package bot

import (
	"bytes"
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
	router.HandleFunc("/bot", FacebookWebHook).Methods("POST")
	router.HandleFunc("/bot", FaceBookVerification).Queries(
		"hub.mode", "{hub.mode}",
		"hub.challenge", "{hub.challenge}",
		"hub.verify_token", "{hub.verify_token}").Methods("GET")
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

	data.MessagingType = "RESPONSE"
	data.Recipient = fbhook.Entry[0].Messaging[0].Sender
	data.Message.Text = fbhook.Entry[0].Messaging[0].Message.Text + " in my ass!"

	sendResponse(data)

	api.WriteJSONResponse(w, data)
}

// FaceBookVerification for testing
func FaceBookVerification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	challenge := r.FormValue("hub.challenge")

	fmt.Fprint(w, challenge)
}

// sendResponse sends response to user or returns error
func sendResponse(ret ReturnStruct) {
	url := "https://graph.facebook.com/v2.6/me/messages?access_token=EAAdVwoaYHFgBAEUB6PRttLApccGXpgVxnXYZA3ZBb6r7ijRjkMfxL2B8sCZC6d4kicG5pqocZCZBtVHGxBUxy4qxv1cSn2bt6ZAFyvn4iFagMwpest5YOkFWma0UC1b69rHE19PlpswRipZBXcXA484Tp6Qg1BDasfP4zwvuUjo1wZDZD"

	data := new(bytes.Buffer)
	var result interface{}
	var err error
	var Client = &http.Client{}

	json.NewEncoder(data).Encode(ret)

	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		fmt.Print(err)
		return
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("messaging_type", "RESPONSE")

	res, err := Client.Do(req)
	if err != nil {
		fmt.Print(err)
		return
	}

	json.NewDecoder(res.Body).Decode(&result)

	fmt.Print(result)
}
