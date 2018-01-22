package handler

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/nickysemenza/food/backend/app/config"
	"github.com/nickysemenza/food/backend/app/model"
	"github.com/nickysemenza/food/backend/app/utils"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

//GetAllRecipes gets all recipes: GET /recipes
func GetAllRecipes(e *config.Env, w http.ResponseWriter, r *http.Request) error {
	var recipes []model.Recipe
	e.DB.Preload("Images").Preload("Categories").Find(&recipes)
	respondSuccess(w, recipes)
	return nil
}

//ErrorTest test todo: remove
func ErrorTest(e *config.Env, w http.ResponseWriter, r *http.Request) error {
	return StatusError{Code: 201, Err: errors.New("sad..")}
}

//GetRecipe gets a recipe by its slug: GET /recipes/{slug}
func GetRecipe(e *config.Env, w http.ResponseWriter, r *http.Request) error {
	recipe := model.Recipe{}
	slug := mux.Vars(r)["slug"]

	if err := recipe.GetFromSlug(e.DB, slug); err != nil {
		return StatusError{Code: 404, Err: errors.New("recipe " + slug + " not found")}
	}

	respondSuccess(w, recipe)
	return nil
}

//PutRecipe updates or creates: PUT /recipes/{slug}
func PutRecipe(e *config.Env, w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	var updatedRecipe model.Recipe
	err := decoder.Decode(&updatedRecipe)
	if err != nil {
		log.Println(err)
	}

	updatedRecipe.CreateOrUpdate(e.DB, false)

	slug := updatedRecipe.Slug
	recipe := model.Recipe{}

	if err := recipe.GetFromSlug(e.DB, slug); err != nil {
		return StatusError{Code: 404, Err: errors.New("recipe " + slug + " not found")}
	}

	respondSuccess(w, recipe)
	return nil
}

//CreateRecipe Creates a new recipe from a Slug and Title
func CreateRecipe(e *config.Env, w http.ResponseWriter, r *http.Request) error {
	//decode the data from JSON encoded request body
	decoder := json.NewDecoder(r.Body)
	var parsed struct {
		Slug  string `json:"slug"`
		Title string `json:"title"`
	}
	err := decoder.Decode(&parsed)
	if err != nil {
		log.Println(err)
	}

	//see if one exists
	recipe := model.Recipe{}
	if !e.DB.Where("slug = ?", parsed.Slug).First(&recipe).RecordNotFound() {
		respondError(w, 500, "slug exists already")
		return nil
	}
	recipe.Slug = parsed.Slug
	recipe.Title = parsed.Title
	e.DB.Save(&recipe)
	respondSuccess(w, "added!")
	return nil
}

//AddNote adds a Note to a Recipe based on Slug, and Note Body
func AddNote(e *config.Env, w http.ResponseWriter, r *http.Request) error {
	//find the recipe we are adding a note to
	recipe := model.Recipe{}
	slug := mux.Vars(r)["slug"]
	if err := e.DB.Where("slug = ?", slug).First(&recipe).Error; err != nil {
		return StatusError{Code: 404, Err: errors.New("recipe " + slug + " not found")}
	}

	//decode the note from JSON encoded request body
	decoder := json.NewDecoder(r.Body)
	var parsed struct {
		Note string `json:"note"`
	}
	err := decoder.Decode(&parsed)
	if err != nil {
		log.Println(err)
	}
	//add a new RecipeNote Model, save it
	note := model.RecipeNote{
		Body:     parsed.Note,
		RecipeID: recipe.ID,
	}
	e.DB.Save(&note)

	respondSuccess(w, note)
	return nil
}

//PutImageUpload uploads images to a recipe based on its Slug
func PutImageUpload(e *config.Env, w http.ResponseWriter, r *http.Request) error {

	var finishedImages []model.Image
	err := r.ParseMultipartForm(100000)
	if err != nil {
		return StatusError{Code: 500, Err: err}
	}

	//get a ref to the parsed multipart form
	m := r.MultipartForm
	slug := m.Value["slug"][0]
	recipe := model.Recipe{}
	if err := e.DB.Where("slug = ?", slug).First(&recipe).Error; err != nil {
		return StatusError{Code: 404, Err: errors.New("recipe " + slug + " not found")}
	}

	files := m.File["file"]
	log.Printf("recieving %d images via upload for recipe %s", len(files), slug)
	for i := range files {
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			return StatusError{Code: 500, Err: err}
		}
		originalFileName := files[i].Filename

		fileData, md5Hash, err := utils.ReadAndHash(file)
		if err != nil {
			return StatusError{Code: 500, Err: err}
		}
		//todo: dedup using md5Hash

		//persist an image obj to DB so we get an PK for s3 path
		imageObj := model.Image{}
		imageObj.Md5Hash = md5Hash
		e.DB.Create(&imageObj)
		e.DB.Model(&recipe).Association("Images").Append(&imageObj)

		//form filesystem / s3 path
		imagePath := fmt.Sprintf("images/%d%s", imageObj.ID, path.Ext(originalFileName))

		os.MkdirAll("public/images", 0777)
		localImageFile, err := os.Create("public/" + imagePath)
		log.Printf("file: %s -> %s", originalFileName, localImageFile.Name())

		defer localImageFile.Close()
		if err != nil {
			return StatusError{Code: 500, Err: err}
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(localImageFile, fileData); err != nil {
			return StatusError{Code: 500, Err: err}
		}

		if os.Getenv("S3_IMAGES") == "true" {

			if err := utils.AddFileToS3(localImageFile.Name(), imagePath); err != nil {
				imageObj.IsInS3 = false
				log.Println(err)
			} else {
				imageObj.IsInS3 = true
				finishedImages = append(finishedImages, imageObj)
			}
		} else {
			finishedImages = append(finishedImages, imageObj)
		}
		imageObj.OriginalFileName = originalFileName
		imageObj.Path = imagePath
		e.DB.Save(&imageObj)
	}
	respondSuccess(w, finishedImages)
	return nil
}

//GetAllImages gets all images, with their related recipes
func GetAllImages(e *config.Env, w http.ResponseWriter, r *http.Request) error {
	var images []model.Image
	e.DB.Preload("Recipes").Find(&images)
	respondSuccess(w, images)
	return nil
}

//GetAllMeals gets all meals, with their related recipes
func GetAllMeals(e *config.Env, w http.ResponseWriter, r *http.Request) error {
	var meals []model.Meal
	e.DB.Order("time DESC").Preload("RecipeMeal.Recipe").Find(&meals)
	respondSuccess(w, meals)
	return nil
}

//GetMealByID retrieves a meal
func GetMealByID(e *config.Env, w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	var meal model.Meal
	e.DB.Preload("RecipeMeal.Recipe").First(&meal, id)
	respondSuccess(w, meal)
	return nil
}

func UpdateMealByID(e *config.Env, w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	decoder := json.NewDecoder(r.Body)
	var updatedMeal model.Meal
	err := decoder.Decode(&updatedMeal)
	if err != nil {
		log.Println(err)
	}

	updatedMeal.CreateOrUpdate(e.DB)

	e.DB.Preload("RecipeMeal.Recipe").First(&updatedMeal, id)
	respondSuccess(w, updatedMeal)
	return nil
}

//GetAllCategories gets all categories that exist
func GetAllCategories(e *config.Env, w http.ResponseWriter, r *http.Request) error {
	var categories []model.Category
	e.DB.Find(&categories)
	respondSuccess(w, categories)
	return nil
}
