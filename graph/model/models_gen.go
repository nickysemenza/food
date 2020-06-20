// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type IngredientInfo interface {
	IsIngredientInfo()
}

type FoodNutrientDerivation struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type NewRecipe struct {
	Name string `json:"name"`
}

type RecipeInput struct {
	UUID         string          `json:"uuid"`
	Name         string          `json:"name"`
	TotalMinutes *int            `json:"totalMinutes"`
	Unit         *string         `json:"unit"`
	Sections     []*SectionInput `json:"sections"`
}

type SectionIngredient struct {
	UUID      string                `json:"uuid"`
	Info      IngredientInfo        `json:"info"`
	Kind      SectionIngredientKind `json:"kind"`
	Grams     float64               `json:"grams"`
	Amount    float64               `json:"amount"`
	Unit      string                `json:"unit"`
	Adjective string                `json:"adjective"`
	Optional  bool                  `json:"optional"`
}

type SectionIngredientInput struct {
	InfoUUID  string                `json:"infoUUID"`
	Kind      SectionIngredientKind `json:"kind"`
	Grams     float64               `json:"grams"`
	Amount    *float64              `json:"amount"`
	Unit      *string               `json:"unit"`
	Adjective *string               `json:"adjective"`
	Optional  *bool                 `json:"optional"`
}

type SectionInput struct {
	Minutes      int                        `json:"minutes"`
	Instructions []*SectionInstructionInput `json:"instructions"`
	Ingredients  []*SectionIngredientInput  `json:"ingredients"`
}

type SectionInstruction struct {
	UUID        string `json:"uuid"`
	Instruction string `json:"instruction"`
}

type SectionInstructionInput struct {
	Instruction string `json:"instruction"`
}

type FoodDataType string

const (
	FoodDataTypeFoundationFood          FoodDataType = "foundation_food"
	FoodDataTypeSampleFood              FoodDataType = "sample_food"
	FoodDataTypeMarketAcquisition       FoodDataType = "market_acquisition"
	FoodDataTypeSurveyFnddsFood         FoodDataType = "survey_fndds_food"
	FoodDataTypeSubSampleFood           FoodDataType = "sub_sample_food"
	FoodDataTypeAgriculturalAcquisition FoodDataType = "agricultural_acquisition"
	FoodDataTypeSrLegacyFood            FoodDataType = "sr_legacy_food"
	FoodDataTypeBrandedFood             FoodDataType = "branded_food"
)

var AllFoodDataType = []FoodDataType{
	FoodDataTypeFoundationFood,
	FoodDataTypeSampleFood,
	FoodDataTypeMarketAcquisition,
	FoodDataTypeSurveyFnddsFood,
	FoodDataTypeSubSampleFood,
	FoodDataTypeAgriculturalAcquisition,
	FoodDataTypeSrLegacyFood,
	FoodDataTypeBrandedFood,
}

func (e FoodDataType) IsValid() bool {
	switch e {
	case FoodDataTypeFoundationFood, FoodDataTypeSampleFood, FoodDataTypeMarketAcquisition, FoodDataTypeSurveyFnddsFood, FoodDataTypeSubSampleFood, FoodDataTypeAgriculturalAcquisition, FoodDataTypeSrLegacyFood, FoodDataTypeBrandedFood:
		return true
	}
	return false
}

func (e FoodDataType) String() string {
	return string(e)
}

func (e *FoodDataType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FoodDataType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FoodDataType", str)
	}
	return nil
}

func (e FoodDataType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SectionIngredientKind string

const (
	SectionIngredientKindRecipe     SectionIngredientKind = "recipe"
	SectionIngredientKindIngredient SectionIngredientKind = "ingredient"
)

var AllSectionIngredientKind = []SectionIngredientKind{
	SectionIngredientKindRecipe,
	SectionIngredientKindIngredient,
}

func (e SectionIngredientKind) IsValid() bool {
	switch e {
	case SectionIngredientKindRecipe, SectionIngredientKindIngredient:
		return true
	}
	return false
}

func (e SectionIngredientKind) String() string {
	return string(e)
}

func (e *SectionIngredientKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SectionIngredientKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SectionIngredientKind", str)
	}
	return nil
}

func (e SectionIngredientKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
