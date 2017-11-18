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

	LoadRecipes()
	CreateRoutes()
	// recipe := FindRecipe("hello")
	// Route(recipe)

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

	err = database.C("messages").Insert(fbhook)
	if err != nil {
		api.ErrorWithJSON(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// Parsing response
	message := DialogFlowQuery{
		Language:  "en",
		Message:   fbhook.Entry[0].Messaging[0].Message.Text,
		SessionID: fbhook.Entry[0].Messaging[0].Sender.ID,
	}

	query, err := parseQuery(message)
	if err != nil {
		api.ErrorWithJSON(w, "Internal error", http.StatusInternalServerError)
		return
	}

	user, err := FBGetUser(query.SessionID)
	if err != nil {
		api.ErrorWithJSON(w, "Internal error", http.StatusInternalServerError)
		return
	}

	var data FBReturnStruct

	data.MessagingType = "RESPONSE"
	data.Recipient = fbhook.Entry[0].Messaging[0].Sender

	recipeName := "404"
	if query.Result.Score > 0 {
		recipeName = query.Result.Metadata.IntentName
	}

	recipe := FindRecipe(recipeName)
	if recipe.Name == "" {
		data.Message.Text = "No recipe for this"
	} else {
		routesData := RouterData{
			Data: map[string]string{
				"username": user.FirstName,
			},
		}
		for k, v := range query.Result.Parameters {
			routesData.Data[k] = v
		}
		fmt.Println(routesData)
		msg := Route(recipe, routesData)
		data.Message.Text = msg.Message
	}

	sendResponse(data)

	api.WriteJSONResponse(w, data)
}

// parseQuery parses messages to parameters
func parseQuery(query DialogFlowQuery) (DialogFlowResponse, error) {
	data := new(bytes.Buffer)
	var result DialogFlowResponse
	var err error
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

// FaceBookVerification for testing
func FaceBookVerification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	challenge := r.FormValue("hub.challenge")

	fmt.Fprint(w, challenge)
}

// sendResponse sends response to user or returns error
func sendResponse(ret FBReturnStruct) {
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
}

// FBGetUser gets a user based on an id
func FBGetUser(id string) (FBUser, error) {
	url := "https://graph.facebook.com/v2.6/" + id

	var result FBUser
	var err error
	var Client = &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer EAAdVwoaYHFgBAEUB6PRttLApccGXpgVxnXYZA3ZBb6r7ijRjkMfxL2B8sCZC6d4kicG5pqocZCZBtVHGxBUxy4qxv1cSn2bt6ZAFyvn4iFagMwpest5YOkFWma0UC1b69rHE19PlpswRipZBXcXA484Tp6Qg1BDasfP4zwvuUjo1wZDZD")

	res, err := Client.Do(req)
	if err != nil {
		return result, err
	}

	json.NewDecoder(res.Body).Decode(&result)

	return result, err
}
