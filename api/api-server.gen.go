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
	// Search foods
	// (GET /foods/search)
	SearchFoods(ctx echo.Context, params SearchFoodsParams) error
	// get a FDC entry by id
	// (GET /foods/{fdc_id})
	GetFoodById(ctx echo.Context, fdcId int) error
	// List all ingredients
	// (GET /ingredients)
	ListIngredients(ctx echo.Context, params ListIngredientsParams) error
	// Create a ingredient
	// (POST /ingredients)
	CreateIngredients(ctx echo.Context) error
	// Converts an ingredient to a recipe, updating all recipes depending on it.
	// (POST /ingredients/{ingredient_id}/convert_to_recipe)
	ConvertIngredientToRecipe(ctx echo.Context, ingredientId string) error
	// Merges the provide ingredients in the body into the param
	// (POST /ingredients/{ingredient_id}/merge)
	MergeIngredients(ctx echo.Context, ingredientId string) error
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

// SearchFoods converts echo context to params.
func (w *ServerInterfaceWrapper) SearchFoods(ctx echo.Context) error {
	var err error

	ctx.Set("bearerAuth.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchFoodsParams
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
	err = w.Handler.SearchFoods(ctx, params)
	return err
}

// GetFoodById converts echo context to params.
func (w *ServerInterfaceWrapper) GetFoodById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "fdc_id" -------------
	var fdcId int

	err = runtime.BindStyledParameter("simple", false, "fdc_id", ctx.Param("fdc_id"), &fdcId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fdc_id: %s", err))
	}

	ctx.Set("bearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetFoodById(ctx, fdcId)
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

// ConvertIngredientToRecipe converts echo context to params.
func (w *ServerInterfaceWrapper) ConvertIngredientToRecipe(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "ingredient_id" -------------
	var ingredientId string

	err = runtime.BindStyledParameter("simple", false, "ingredient_id", ctx.Param("ingredient_id"), &ingredientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter ingredient_id: %s", err))
	}

	ctx.Set("bearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ConvertIngredientToRecipe(ctx, ingredientId)
	return err
}

// MergeIngredients converts echo context to params.
func (w *ServerInterfaceWrapper) MergeIngredients(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "ingredient_id" -------------
	var ingredientId string

	err = runtime.BindStyledParameter("simple", false, "ingredient_id", ctx.Param("ingredient_id"), &ingredientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter ingredient_id: %s", err))
	}

	ctx.Set("bearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.MergeIngredients(ctx, ingredientId)
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
	router.GET("/foods/search", wrapper.SearchFoods)
	router.GET("/foods/:fdc_id", wrapper.GetFoodById)
	router.GET("/ingredients", wrapper.ListIngredients)
	router.POST("/ingredients", wrapper.CreateIngredients)
	router.POST("/ingredients/:ingredient_id/convert_to_recipe", wrapper.ConvertIngredientToRecipe)
	router.POST("/ingredients/:ingredient_id/merge", wrapper.MergeIngredients)
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
