package bot

import (
	"encoding/json"
	"net/http"

	"github.com/cloud-oblig3/api"
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
	router.HandleFunc("/bot/", FacebookWebHook).Methods("POST")

}

// FacebookWebHook handler for Facebook Webook events
func FacebookWebHook(w http.ResponseWriter, r *http.Request) {
	var fbhook FBWebHook
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&fbhook)

	if err != nil {
		api.ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
		return
	}

	err2 := database.C("messages").Insert(fbhook)
	if err2 != nil {
		api.ErrorWithJSON(w, "Internal error", http.StatusInternalServerError)
		return
	}

	api.WriteJSONResponse(w, fbhook)
}
