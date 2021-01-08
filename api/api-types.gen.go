// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"time"
)

// AuthResp defines model for AuthResp.
type AuthResp struct {
	Jwt  string                 `json:"jwt"`
	User map[string]interface{} `json:"user"`
}

// BrandedFood defines model for BrandedFood.
type BrandedFood struct {
	BrandOwner          *string `json:"brand_owner,omitempty"`
	BrandedFoodCategory *string `json:"branded_food_category,omitempty"`
	HouseholdServing    *string `json:"household_serving,omitempty"`
	Ingredients         *string `json:"ingredients,omitempty"`
	ServingSize         float64 `json:"serving_size"`
	ServingSizeUnit     string  `json:"serving_size_unit"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Food defines model for Food.
type Food struct {

	// branded_food
	BrandedInfo *BrandedFood `json:"branded_info,omitempty"`

	// food category, set for some
	Category *FoodCategory `json:"category,omitempty"`

	// todo
	DataType string `json:"data_type"`

	// Food description
	Description string `json:"description"`

	// FDC Id
	FdcId int `json:"fdc_id"`

	// todo
	Nutrients []FoodNutrient `json:"nutrients"`

	// portion datapoints
	Portions *[]FoodPortion `json:"portions,omitempty"`
}

// FoodCategory defines model for FoodCategory.
type FoodCategory struct {

	// Food description
	Code string `json:"code"`

	// Food description
	Description string `json:"description"`
}

// FoodNutrient defines model for FoodNutrient.
type FoodNutrient struct {
	Amount     float64 `json:"amount"`
	DataPoints int     `json:"data_points"`

	// todo
	Nutrient Nutrient `json:"nutrient"`
}

// FoodPortion defines model for FoodPortion.
type FoodPortion struct {
	Amount             float64 `json:"amount"`
	GramWeight         float64 `json:"gram_weight"`
	Id                 int     `json:"id"`
	Modifier           string  `json:"modifier"`
	PortionDescription string  `json:"portion_description"`
}

// GooglePhoto defines model for GooglePhoto.
type GooglePhoto struct {

	// public image
	BaseUrl string `json:"base_url"`

	// blur hash
	BlurHash *string `json:"blur_hash,omitempty"`

	// when it was taken
	Created time.Time `json:"created"`

	// height px
	Height int64 `json:"height"`

	// id
	Id string `json:"id"`

	// width px
	Width int64 `json:"width"`
}

// GooglePhotosAlbum defines model for GooglePhotosAlbum.
type GooglePhotosAlbum struct {

	// id
	Id string `json:"id"`

	// product_url
	ProductUrl string `json:"product_url"`

	// title
	Title string `json:"title"`

	// usecase
	Usecase string `json:"usecase"`
}

// Ingredient defines model for Ingredient.
type Ingredient struct {

	// id
	Id string `json:"id"`

	// Ingredient name
	Name string `json:"name"`
}

// IngredientDetail defines model for IngredientDetail.
type IngredientDetail struct {

	// Ingredients that are equivalent
	Children *[]IngredientDetail `json:"children,omitempty"`

	// An Ingredient
	Ingredient Ingredient `json:"ingredient"`

	// Recipes referencing this ingredient
	Recipes *[]RecipeDetail `json:"recipes,omitempty"`
}

// List defines model for List.
type List struct {

	// How many items were requested for this page
	Limit int `json:"limit"`

	// todo
	Offset int `json:"offset"`

	// Total number of pages available
	PageCount int `json:"page_count"`

	// What number page this is
	PageNumber int `json:"page_number"`

	// Total number of items across all pages
	TotalCount int `json:"total_count"`
}

// Meal defines model for Meal.
type Meal struct {

	// when it was taken
	AteAt time.Time `json:"ate_at"`

	// id
	Id string `json:"id"`

	// public image
	Name    string        `json:"name"`
	Photos  []GooglePhoto `json:"photos"`
	Recipes *[]MealRecipe `json:"recipes,omitempty"`
}

// MealRecipe defines model for MealRecipe.
type MealRecipe struct {

	// when it was taken
	Multiplier float64 `json:"multiplier"`

	// A revision of a recipe
	Recipe RecipeDetail `json:"recipe"`
}

// MealRecipeUpdate defines model for MealRecipeUpdate.
type MealRecipeUpdate struct {

	// todo
	Action string `json:"action"`

	// multiplier
	Multiplier float64 `json:"multiplier"`

	// Recipe Id
	RecipeId string `json:"recipe_id"`
}

// Nutrient defines model for Nutrient.
type Nutrient struct {

	// todo
	Id int `json:"id"`

	// todo
	Name string `json:"name"`

	// todo
	UnitName string `json:"unit_name"`
}

// PaginatedFoods defines model for PaginatedFoods.
type PaginatedFoods struct {
	Foods *[]Food `json:"foods,omitempty"`

	// A generic list (for pagination use)
	Meta *List `json:"meta,omitempty"`
}

// PaginatedIngredients defines model for PaginatedIngredients.
type PaginatedIngredients struct {
	Ingredients *[]IngredientDetail `json:"ingredients,omitempty"`

	// A generic list (for pagination use)
	Meta *List `json:"meta,omitempty"`
}

// PaginatedMeals defines model for PaginatedMeals.
type PaginatedMeals struct {
	Meals *[]Meal `json:"meals,omitempty"`

	// A generic list (for pagination use)
	Meta *List `json:"meta,omitempty"`
}

// PaginatedPhotos defines model for PaginatedPhotos.
type PaginatedPhotos struct {

	// A generic list (for pagination use)
	Meta   *List          `json:"meta,omitempty"`
	Photos *[]GooglePhoto `json:"photos,omitempty"`
}

// PaginatedRecipes defines model for PaginatedRecipes.
type PaginatedRecipes struct {

	// A generic list (for pagination use)
	Meta    *List     `json:"meta,omitempty"`
	Recipes *[]Recipe `json:"recipes,omitempty"`
}

// Recipe defines model for Recipe.
type Recipe struct {

	// id
	Id string `json:"id"`

	// all the versions of the recipe
	Versions []RecipeDetail `json:"versions"`
}

// RecipeDetail defines model for RecipeDetail.
type RecipeDetail struct {

	// id
	Id string `json:"id"`

	// whether or not it is the most recent version
	IsLatestVersion *bool `json:"is_latest_version,omitempty"`

	// recipe name
	Name string `json:"name"`

	// serving quantity
	Quantity int64 `json:"quantity"`

	// sections of the recipe
	Sections []RecipeSection `json:"sections"`

	// num servings
	Servings *int64 `json:"servings,omitempty"`

	// book or websites
	Sources *[]RecipeSource `json:"sources,omitempty"`

	// serving unit
	Unit string `json:"unit"`

	// version of the recipe
	Version *int64 `json:"version,omitempty"`
}

// RecipeSection defines model for RecipeSection.
type RecipeSection struct {

	// A range of time or a specific duration of time (in seconds)
	Duration *TimeRange `json:"duration,omitempty"`

	// id
	Id string `json:"id"`

	// x
	Ingredients []SectionIngredient `json:"ingredients"`

	// x
	Instructions []SectionInstruction `json:"instructions"`
}

// RecipeSource defines model for RecipeSource.
type RecipeSource struct {

	// page number/section (if book)
	Page *string `json:"page,omitempty"`

	// title (if book)
	Title *string `json:"title,omitempty"`

	// url
	Url *string `json:"url,omitempty"`
}

// RecipeWrapper defines model for RecipeWrapper.
type RecipeWrapper struct {

	// A revision of a recipe
	Detail RecipeDetail `json:"detail"`

	// id
	Id string `json:"id"`
}

// SearchResult defines model for SearchResult.
type SearchResult struct {

	// The ingredients
	Ingredients *[]Ingredient `json:"ingredients,omitempty"`

	// A generic list (for pagination use)
	Meta *List `json:"meta,omitempty"`

	// The recipes
	Recipes *[]RecipeWrapper `json:"recipes,omitempty"`
}

// SectionIngredient defines model for SectionIngredient.
type SectionIngredient struct {

	// adjective
	Adjective *string `json:"adjective,omitempty"`

	// amount
	Amount *float64 `json:"amount,omitempty"`

	// weight in grams
	Grams float64 `json:"grams"`

	// id
	Id string `json:"id"`

	// An Ingredient
	Ingredient *Ingredient `json:"ingredient,omitempty"`

	// what kind of ingredient
	Kind string `json:"kind"`

	// optional
	Optional *bool `json:"optional,omitempty"`

	// raw line item (pre-import/scrape)
	Original *string `json:"original,omitempty"`

	// A revision of a recipe
	Recipe *RecipeDetail `json:"recipe,omitempty"`

	// x
	Substitutes *[]SectionIngredient `json:"substitutes,omitempty"`

	// unit
	Unit *string `json:"unit,omitempty"`
}

// SectionInstruction defines model for SectionInstruction.
type SectionInstruction struct {

	// id
	Id string `json:"id"`

	// instruction
	Instruction string `json:"instruction"`
}

// TimeRange defines model for TimeRange.
type TimeRange struct {

	// The maximum amount of seconds (if a range)
	Max int `json:"max"`

	// The minimum amount of seconds (or the total, if not a range)
	Min int `json:"min"`
}

// LimitParam defines model for limitParam.
type LimitParam int

// NameParam defines model for nameParam.
type NameParam string

// OffsetParam defines model for offsetParam.
type OffsetParam int

// AuthLoginParams defines parameters for AuthLogin.
type AuthLoginParams struct {

	// Google code
	Code string `json:"code"`
}

// SearchFoodsParams defines parameters for SearchFoods.
type SearchFoodsParams struct {

	// The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `json:"offset,omitempty"`

	// The numbers of items to return.
	Limit *LimitParam `json:"limit,omitempty"`

	// The search query (name).
	Name NameParam `json:"name"`
}

// ListIngredientsParams defines parameters for ListIngredients.
type ListIngredientsParams struct {

	// The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `json:"offset,omitempty"`

	// The numbers of items to return.
	Limit *LimitParam `json:"limit,omitempty"`
}

// CreateIngredientsJSONBody defines parameters for CreateIngredients.
type CreateIngredientsJSONBody Ingredient

// MergeIngredientsJSONBody defines parameters for MergeIngredients.
type MergeIngredientsJSONBody struct {
	IngredientIds []string `json:"ingredient_ids"`
}

// ListMealsParams defines parameters for ListMeals.
type ListMealsParams struct {

	// The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `json:"offset,omitempty"`

	// The numbers of items to return.
	Limit *LimitParam `json:"limit,omitempty"`
}

// UpdateRecipesForMealJSONBody defines parameters for UpdateRecipesForMeal.
type UpdateRecipesForMealJSONBody MealRecipeUpdate

// ListPhotosParams defines parameters for ListPhotos.
type ListPhotosParams struct {

	// The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `json:"offset,omitempty"`

	// The numbers of items to return.
	Limit *LimitParam `json:"limit,omitempty"`
}

// ListRecipesParams defines parameters for ListRecipes.
type ListRecipesParams struct {

	// The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `json:"offset,omitempty"`

	// The numbers of items to return.
	Limit *LimitParam `json:"limit,omitempty"`
}

// CreateRecipesJSONBody defines parameters for CreateRecipes.
type CreateRecipesJSONBody RecipeWrapper

// SearchParams defines parameters for Search.
type SearchParams struct {

	// The number of items to skip before starting to collect the result set.
	Offset *OffsetParam `json:"offset,omitempty"`

	// The numbers of items to return.
	Limit *LimitParam `json:"limit,omitempty"`

	// The search query (name).
	Name NameParam `json:"name"`
}

// CreateIngredientsRequestBody defines body for CreateIngredients for application/json ContentType.
type CreateIngredientsJSONRequestBody CreateIngredientsJSONBody

// MergeIngredientsRequestBody defines body for MergeIngredients for application/json ContentType.
type MergeIngredientsJSONRequestBody MergeIngredientsJSONBody

// UpdateRecipesForMealRequestBody defines body for UpdateRecipesForMeal for application/json ContentType.
type UpdateRecipesForMealJSONRequestBody UpdateRecipesForMealJSONBody

// CreateRecipesRequestBody defines body for CreateRecipes for application/json ContentType.
type CreateRecipesJSONRequestBody CreateRecipesJSONBody
