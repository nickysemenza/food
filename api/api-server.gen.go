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
	// Assosiates a food with a given ingredient
	// (POST /ingredients/{ingredient_id}/associate_food)
	AssociateFoodWithIngredient(ctx echo.Context, ingredientId string, params AssociateFoodWithIngredientParams) error
	// Converts an ingredient to a recipe, updating all recipes depending on it.
	// (POST /ingredients/{ingredient_id}/convert_to_recipe)
	ConvertIngredientToRecipe(ctx echo.Context, ingredientId string) error
	// Merges the provide ingredients in the body into the param
	// (POST /ingredients/{ingredient_id}/merge)
	MergeIngredients(ctx echo.Context, ingredientId string) error
	// List all meals
	// (GET /meals)
	ListMeals(ctx echo.Context, params ListMealsParams) error
	// Info for a specific meal
	// (GET /meals/{meal_id})
	GetMealById(ctx echo.Context, mealId string) error
	// Update the recipes associated with a given meal
	// (PATCH /meals/{meal_id}/recipes)
	UpdateRecipesForMeal(ctx echo.Context, mealId string) error
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

// AssociateFoodWithIngredient converts echo context to params.
func (w *ServerInterfaceWrapper) AssociateFoodWithIngredient(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "ingredient_id" -------------
	var ingredientId string

	err = runtime.BindStyledParameter("simple", false, "ingredient_id", ctx.Param("ingredient_id"), &ingredientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter ingredient_id: %s", err))
	}

	ctx.Set("bearerAuth.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params AssociateFoodWithIngredientParams
	// ------------- Required query parameter "fdc_id" -------------

	err = runtime.BindQueryParameter("form", true, true, "fdc_id", ctx.QueryParams(), &params.FdcId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fdc_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AssociateFoodWithIngredient(ctx, ingredientId, params)
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

// GetMealById converts echo context to params.
func (w *ServerInterfaceWrapper) GetMealById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "meal_id" -------------
	var mealId string

	err = runtime.BindStyledParameter("simple", false, "meal_id", ctx.Param("meal_id"), &mealId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter meal_id: %s", err))
	}

	ctx.Set("bearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetMealById(ctx, mealId)
	return err
}

// UpdateRecipesForMeal converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateRecipesForMeal(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "meal_id" -------------
	var mealId string

	err = runtime.BindStyledParameter("simple", false, "meal_id", ctx.Param("meal_id"), &mealId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter meal_id: %s", err))
	}

	ctx.Set("bearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateRecipesForMeal(ctx, mealId)
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
	router.POST("/ingredients/:ingredient_id/associate_food", wrapper.AssociateFoodWithIngredient)
	router.POST("/ingredients/:ingredient_id/convert_to_recipe", wrapper.ConvertIngredientToRecipe)
	router.POST("/ingredients/:ingredient_id/merge", wrapper.MergeIngredients)
	router.GET("/meals", wrapper.ListMeals)
	router.GET("/meals/:meal_id", wrapper.GetMealById)
	router.PATCH("/meals/:meal_id/recipes", wrapper.UpdateRecipesForMeal)
	router.GET("/photos", wrapper.ListPhotos)
	router.GET("/recipes", wrapper.ListRecipes)
	router.POST("/recipes", wrapper.CreateRecipes)
	router.GET("/recipes/:recipe_id", wrapper.GetRecipeById)
	router.GET("/search", wrapper.Search)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"278jGfjTg37/67v/BAAA//9iQBEaVlgAAA==",
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
