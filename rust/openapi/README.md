# Rust API client for openapi

API for https://github.com/nickysemenza/gourd

## Overview

This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project. By using the [openapi-spec](https://openapis.org) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.RustClientCodegen

## Installation

Put the package under your project folder and add the following to `Cargo.toml` under `[dependencies]`:

```
    openapi = { path = "./generated" }
```

## Documentation for API Endpoints

All URIs are relative to _http://localhost:4242/api_

| Class | Method | HTTP request | Description |
| ----- | ------ | ------------ | ----------- |

## Documentation For Models

- [AuthResp](docs/AuthResp.md)
- [BrandedFood](docs/BrandedFood.md)
- [Error](docs/Error.md)
- [Food](docs/Food.md)
- [FoodCategory](docs/FoodCategory.md)
- [FoodDataType](docs/FoodDataType.md)
- [FoodNutrient](docs/FoodNutrient.md)
- [FoodNutrientUnit](docs/FoodNutrientUnit.md)
- [FoodPortion](docs/FoodPortion.md)
- [GooglePhoto](docs/GooglePhoto.md)
- [GooglePhotosAlbum](docs/GooglePhotosAlbum.md)
- [Ingredient](docs/Ingredient.md)
- [IngredientDetail](docs/IngredientDetail.md)
- [InlineObject](docs/InlineObject.md)
- [InlineResponse200](docs/InlineResponse200.md)
- [Items](docs/Items.md)
- [Meal](docs/Meal.md)
- [MealRecipe](docs/MealRecipe.md)
- [MealRecipeUpdate](docs/MealRecipeUpdate.md)
- [Nutrient](docs/Nutrient.md)
- [PaginatedFoods](docs/PaginatedFoods.md)
- [PaginatedIngredients](docs/PaginatedIngredients.md)
- [PaginatedMeals](docs/PaginatedMeals.md)
- [PaginatedPhotos](docs/PaginatedPhotos.md)
- [PaginatedRecipes](docs/PaginatedRecipes.md)
- [Recipe](docs/Recipe.md)
- [RecipeDetail](docs/RecipeDetail.md)
- [RecipeSection](docs/RecipeSection.md)
- [RecipeSource](docs/RecipeSource.md)
- [RecipeWrapper](docs/RecipeWrapper.md)
- [SearchResult](docs/SearchResult.md)
- [SectionIngredient](docs/SectionIngredient.md)
- [SectionInstruction](docs/SectionInstruction.md)
- [TimeRange](docs/TimeRange.md)

To get access to the crate's generated documentation, use:

```
cargo doc --open
```

## Author