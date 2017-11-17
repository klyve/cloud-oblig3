package bot

import (
	"bytes"
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

	// Making and parsing response
	message := DialogFlowQuery{
		Language:  "en",
		Message:   fbhook.Entry[0].Messaging[0].Message.Text,
		SessionID: fbhook.Entry[0].Messaging[0].Sender.ID,
	}

	query, err := parseQuery(message)
	if err != nil {
		fmt.Print(err)
		api.ErrorWithJSON(w, "Internal error", http.StatusInternalServerError)
		return
	}

	var data FBReturnStruct

	data.MessagingType = "Standard Messaging"
	data.Recipient = fbhook.Entry[0].Messaging[0].Recipient
	data.Message.Text = "Whaddup my nigguh!!"

	sendResponse(data)

	api.WriteJSONResponse(w, data)
}

// parseQuery parses messages to parameters
func parseQuery(query DialogFlowQuery) (DialogFlowResponse, error) {
	data := new(bytes.Buffer)
	var result DialogFlowResponse
	var err error = nil
	var Client = &http.Client{}

	json.NewEncoder(data).Encode(query)

	req, err := http.NewRequest("POST", "https://api.dialogflow.com/v1/query", data)
	if err != nil {
		return result, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer 8f67b21441414cf7853e8d850789cb5f")

	res, err := Client.Do(req)
	if err != nil {
		return result, err
	}

	json.NewDecoder(res.Body).Decode(&result)

	return result, err
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
