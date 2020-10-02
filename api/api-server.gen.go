// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List all ingredients
	// (GET /ingredients)
	ListIngredients(ctx echo.Context, params ListIngredientsParams) error
	// Create a ingredient
	// (POST /ingredients)
	CreateIngredients(ctx echo.Context) error
	// List all photos
	// (GET /photos)
	ListPhotos(ctx echo.Context, params ListPhotosParams) error
	// List all recipes
	// (GET /recipes)
	ListRecipes(ctx echo.Context, params ListRecipesParams) error
	// Create a recipe
	// (POST /recipes)
	CreateRecipes(ctx echo.Context) error
	// Info for a specific recipe
	// (GET /recipes/{recipe_id})
	GetRecipeById(ctx echo.Context, recipeId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ListIngredients converts echo context to params.
func (w *ServerInterfaceWrapper) ListIngredients(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListIngredientsParams
	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListIngredients(ctx, params)
	return err
}

// CreateIngredients converts echo context to params.
func (w *ServerInterfaceWrapper) CreateIngredients(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateIngredients(ctx)
	return err
}

// ListPhotos converts echo context to params.
func (w *ServerInterfaceWrapper) ListPhotos(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListPhotosParams
	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListPhotos(ctx, params)
	return err
}

// ListRecipes converts echo context to params.
func (w *ServerInterfaceWrapper) ListRecipes(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListRecipesParams
	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListRecipes(ctx, params)
	return err
}

// CreateRecipes converts echo context to params.
func (w *ServerInterfaceWrapper) CreateRecipes(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateRecipes(ctx)
	return err
}

// GetRecipeById converts echo context to params.
func (w *ServerInterfaceWrapper) GetRecipeById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "recipe_id" -------------
	var recipeId string

	err = runtime.BindStyledParameter("simple", false, "recipe_id", ctx.Param("recipe_id"), &recipeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter recipe_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRecipeById(ctx, recipeId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/ingredients", wrapper.ListIngredients)
	router.POST("/ingredients", wrapper.CreateIngredients)
	router.GET("/photos", wrapper.ListPhotos)
	router.GET("/recipes", wrapper.ListRecipes)
	router.POST("/recipes", wrapper.CreateRecipes)
	router.GET("/recipes/:recipe_id", wrapper.GetRecipeById)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
