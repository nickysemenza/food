package app

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	h "github.com/nickysemenza/food/backend/app/handler"
	"github.com/nickysemenza/food/backend/app/model"
	"log"
	"net/http"
)

type App struct {
	R *mux.Router
}
type Route struct {
	Method      string
	Pattern     string
	HandlerFunc func(e *h.Env, w http.ResponseWriter, r *http.Request) error
}

type Routes []Route

func (a *App) Initialize(config *Config) *h.Env {
	db, err := gorm.Open(config.DB.Dialect, config.getDBURI())
	if err != nil {
		log.Fatal("Could not connect database")
	}
	//set up the env
	env := &h.Env{
		DB:     model.DBMigrate(db),
		Port:   config.Port,
		Router: &a.R,
	}
	a.buildRoutes(env)
	return env
}

func (a *App) buildRoutes(env *h.Env) {

	var routes = Routes{
		{"GET", "/", h.ErrorTest},
		{"PUT", "/imageupload", h.ImageUploadTest},
		{"GET", "/recipes", h.GetAllRecipes},
		{"POST", "/recipes", h.CreateRecipe},
		{"GET", "/recipes/{slug}", h.GetRecipe},
		{"PUT", "/recipes/{slug}", h.PutRecipe},
		{"POST", "/recipes/{slug}/notes", h.AddNote},
		{"GET", "/images", h.GetAllImages},
		{"GET", "/categories", h.GetAllCategories},
		{"GET", "/meals", h.GetAllMeals},
	}

	//add them all
	a.R = mux.NewRouter()
	for _, route := range routes {
		a.R.Handle(route.Pattern, h.Handler{env, route.HandlerFunc}).Methods(route.Method)
	}
	a.R.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	a.R.NotFoundHandler = h.Handler{env, h.NotFoundRoute}
}

func (a *App) RunServer(host string) {
	log.Println("Running API server on", host)
	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(host, handlers.CORS(originsOk, headersOk, methodsOk)(a.R)))
}
