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
	// Get foods
	// (GET /foods/bulk)
	GetFoodsByIds(ctx echo.Context, params GetFoodsByIdsParams) error
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
	// Get a specific ingredient
	// (GET /ingredients/{ingredient_id})
	GetIngredientById(ctx echo.Context, ingredientId string) error
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

// GetFoodsByIds converts echo context to params.
func (w *ServerInterfaceWrapper) GetFoodsByIds(ctx echo.Context) error {
	var err error

	ctx.Set("bearerAuth.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetFoodsByIdsParams
	// ------------- Required query parameter "fdc_id" -------------

	err = runtime.BindQueryParameter("form", true, true, "fdc_id", ctx.QueryParams(), &params.FdcId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fdc_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetFoodsByIds(ctx, params)
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

	// ------------- Optional query parameter "data_types" -------------

	err = runtime.BindQueryParameter("form", true, false, "data_types", ctx.QueryParams(), &params.DataTypes)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter data_types: %s", err))
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

// GetIngredientById converts echo context to params.
func (w *ServerInterfaceWrapper) GetIngredientById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "ingredient_id" -------------
	var ingredientId string

	err = runtime.BindStyledParameter("simple", false, "ingredient_id", ctx.Param("ingredient_id"), &ingredientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter ingredient_id: %s", err))
	}

	ctx.Set("bearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetIngredientById(ctx, ingredientId)
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
	router.GET("/foods/bulk", wrapper.GetFoodsByIds)
	router.GET("/foods/search", wrapper.SearchFoods)
	router.GET("/foods/:fdc_id", wrapper.GetFoodById)
	router.GET("/ingredients", wrapper.ListIngredients)
	router.POST("/ingredients", wrapper.CreateIngredients)
	router.GET("/ingredients/:ingredient_id", wrapper.GetIngredientById)
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
	"///gBMCmrF8AAA==",
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
