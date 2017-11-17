package bot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cloud-oblig3/api"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

var database *mgo.Database

// Init bot
func Init(router *mux.Router, db *mgo.Database) {
	database = db
	router.HandleFunc("/bot/", FacebookWebHook).Methods("POST")
	router.HandleFunc("/bot/", HelloBot).Methods("GET")
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

	var data ReturnStruct
	data.MessagingType = "Standard Messaging"
	data.Recipient = fbhook.Entry[0].Messaging[0].Recipient
	data.Message.Text = "Whaddup my nigguh!!"

	api.WriteJSONResponse(w, data)
}

// HelloBot for testing
func HelloBot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	fmt.Fprintln(w, "Hello world")
}
