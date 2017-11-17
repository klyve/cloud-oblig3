package bot

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

var database *mgo.Database

// Init bot
func Init(router *mux.Router, db *mgo.Database) {
	database = db
	// router.HandleFunc("/exchange/", GetExchange).Methods("GET")
	LoadRecipes()
	CreateRoutes()
	recipe := FindRecipe("hello")
	Route(recipe)
}
