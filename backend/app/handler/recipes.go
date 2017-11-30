package handler

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/nickysemenza/food/backend/app/model"
	"net/http"
)

func GetAllRecipes(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	recipes := []model.Recipe{}
	db.Select([]string{"slug"}).Find(&recipes)
	var slugs []string
	for _, r := range recipes {
		slugs = append(slugs, r.Slug)
	}
	respondJSON(w, http.StatusOK, slugs)
}
func GetRecipe(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	recipe := model.Recipe{}
	vars := mux.Vars(r)
	slug := vars["slug"]
	db.Where("slug = ?", slug).Preload("Sections.Instructions").Preload("Sections.Ingredients.Item").First(&recipe)
	respondJSON(w, http.StatusOK, recipe)
}
