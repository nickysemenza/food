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
	// List all albums
	// (GET /albums)
	ListAllAlbums(ctx echo.Context) error
	// Google Login callback
	// (POST /auth)
	AuthLogin(ctx echo.Context, params AuthLoginParams) error
	// List all ingredients
	// (GET /ingredients)
	ListIngredients(ctx echo.Context, params ListIngredientsParams) error
	// Create a ingredient
	// (POST /ingredients)
	CreateIngredients(ctx echo.Context) error
	// List all meals
	// (GET /meals)
	ListMeals(ctx echo.Context, params ListMealsParams) error
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
	// Search recipes and ingredients
	// (GET /search)
	Search(ctx echo.Context, params SearchParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ListAllAlbums converts echo context to params.
func (w *ServerInterfaceWrapper) ListAllAlbums(ctx echo.Context) error {
	var err error

	ctx.Set("bearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListAllAlbums(ctx)
	return err
}

// AuthLogin converts echo context to params.
func (w *ServerInterfaceWrapper) AuthLogin(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params AuthLoginParams
	// ------------- Required query parameter "code" -------------

	err = runtime.BindQueryParameter("form", true, true, "code", ctx.QueryParams(), &params.Code)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter code: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AuthLogin(ctx, params)
	return err
}

// ListIngredients converts echo context to params.
func (w *ServerInterfaceWrapper) ListIngredients(ctx echo.Context) error {
	var err error

	ctx.Set("bearerAuth.Scopes", []string{""})

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

	ctx.Set("bearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateIngredients(ctx)
	return err
}

// ListMeals converts echo context to params.
func (w *ServerInterfaceWrapper) ListMeals(ctx echo.Context) error {
	var err error

	ctx.Set("bearerAuth.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListMealsParams
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
	err = w.Handler.ListMeals(ctx, params)
	return err
}

// ListPhotos converts echo context to params.
func (w *ServerInterfaceWrapper) ListPhotos(ctx echo.Context) error {
	var err error

	ctx.Set("bearerAuth.Scopes", []string{""})

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

	ctx.Set("bearerAuth.Scopes", []string{""})

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

	ctx.Set("bearerAuth.Scopes", []string{""})

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

	ctx.Set("bearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRecipeById(ctx, recipeId)
	return err
}

// Search converts echo context to params.
func (w *ServerInterfaceWrapper) Search(ctx echo.Context) error {
	var err error

	ctx.Set("bearerAuth.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchParams
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

	// ------------- Required query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, true, "name", ctx.QueryParams(), &params.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Search(ctx, params)
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

	router.GET("/albums", wrapper.ListAllAlbums)
	router.POST("/auth", wrapper.AuthLogin)
	router.GET("/ingredients", wrapper.ListIngredients)
	router.POST("/ingredients", wrapper.CreateIngredients)
	router.GET("/meals", wrapper.ListMeals)
	router.GET("/photos", wrapper.ListPhotos)
	router.GET("/recipes", wrapper.ListRecipes)
	router.POST("/recipes", wrapper.CreateRecipes)
	router.GET("/recipes/:recipe_id", wrapper.GetRecipeById)
	router.GET("/search", wrapper.Search)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"/zLbH4wxqJXM6f5u/78AAAD//9Fd4zzyLwAA",
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
