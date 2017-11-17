package bot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klyve/cloud-oblig3/api"
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
	router.HandleFunc("/bot/", HelloBot).Methods("GET")

	router.HandleFunc("/bot", FacebookWebHook).Methods("POST")
	router.HandleFunc("/bot", HelloBot).Methods("GET")

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

	err = database.C("messages").Insert(fbhook)
	if err != nil {
		api.ErrorWithJSON(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// TEST RESPONSE

	var data FBReturnStruct

	data.MessagingType = "Standard Messaging"
	data.Recipient = fbhook.Entry[0].Messaging[0].Recipient
	data.Message.Text = "Whaddup my nigguh!!"

	sendResponse(data)

	api.WriteJSONResponse(w, data)
}

func parseQuery(query DialogFlowQuery) DialogFlowResponse {

}

// HelloBot for testing
func HelloBot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	fmt.Fprintln(w, "Hello world")
}

// sendResponse sends response to user or returns error
func sendResponse(data FBReturnStruct) {
	url := "https://graph.facebook.com/v2.6/me/messages?access_token=EAAdVwoaYHFgBADZCsmyG5e87NUL6ardZBVEDmFFPPIyZAifF1hMLaKpdqQuwZCcQmI4tCgNvEiGG0bwsPZCbsqZCGNZBphG4N8VZAAtLkRXjPzAdI6KGYpGLFXppwUuZAaBN8ibZCoQWX0eZCDdV9V3ZAMOmEYeALRjbZCWnHJE2GZAm9RsgZDZD"
	output, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return
	}

	http.Post(url, string(output), nil)
}
