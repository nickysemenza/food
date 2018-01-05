package handler

import (
	"encoding/json"
	"github.com/nickysemenza/food/backend/app/config"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
func respondSuccess(w http.ResponseWriter, payload interface{}) {
	respondJSON(w, http.StatusOK, payload)
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func NotFoundRoute(e *config.Env, w http.ResponseWriter, r *http.Request) error {
	return StatusError{Code: 404, Err: errors.New("route not found: " + r.RequestURI)}
}

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// The Handler struct that takes a configured Env and a function matching
// our useful signature.
type Handler struct {
	*config.Env
	H func(e *config.Env, w http.ResponseWriter, r *http.Request) error
	P bool
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"method": r.Method,
		"URI":    r.RequestURI,
	}).Info("http_request")

	if h.P == true {
		authorized := false
		tokFromHeader := r.Header.Get("X-Jwt")
		log.Printf("got token: %s", tokFromHeader)
		if tokFromHeader != "" {
			u, _ := getUserFromToken(h.Env, tokFromHeader)
			if u != nil && u.Admin == true {
				authorized = true
				h.Env.CurrentUser = u
			}
		}

		if !authorized {
			respondError(w, http.StatusUnauthorized, "not authorized")
			return
		}
	}
	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status(), e)
			respondError(w, e.Status(), e.Error())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
