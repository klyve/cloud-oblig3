package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klyve/cloud-oblig2/currency"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var database *mgo.Database

// Init the api
func Init(router *mux.Router, db *mgo.Database) {
	database = db
	router.HandleFunc("/exchange", GetExchange).Methods("GET")
	router.HandleFunc("/exchange", PostExchange).Methods("POST")
	router.HandleFunc("/exchange/latest", GetLatestRates).Methods("POST")
	router.HandleFunc("/exchange/evaluationtrigger", TestHooks).Methods("GET")
	router.HandleFunc("/exchange/average", AverageRates).Methods("POST")

	// DIRTY
	router.HandleFunc("/exchange/average", GetExchange).Methods("GET")
	router.HandleFunc("/exchange/latest", GetExchange).Methods("GET")

	router.HandleFunc("/exchange/{id}", GetWebhookData).Methods("GET")
	router.HandleFunc("/exchange/{id}", DeleteWebhook).Methods("DELETE")
}

// AverageRates api
func AverageRates(w http.ResponseWriter, r *http.Request) {
	var latest LatestRates
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&latest)

	if err != nil {
		ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
		return
	}

	var curr []currency.DataList
	err2 := database.C("currency").Find(nil).Limit(7).Sort("-_id").All(&curr)
	if err2 != nil {
		ErrorWithJSON(w, "Internal error", http.StatusInternalServerError)
		return
	}
	var avg float32
	avg = 0
	for i := range curr {
		item := curr[i]
		avg += item.From(latest.BaseCurrency).To(latest.TargetCurrency).Rate
	}
	fmt.Fprint(w, avg/7)
}

// InvokeWebHook api
func InvokeWebHook(curr currency.Convertion, item WebHook) {
	inv := WebHookInvoked{
		BaseCurrency:    curr.From,
		TargetCurrency:  curr.To,
		CurrentRate:     curr.Rate,
		MinTriggerValue: item.MinTriggerValue,
		MaxTriggerValue: item.MaxTriggerValue,
	}
	data := new(bytes.Buffer)
	json.NewEncoder(data).Encode(inv)
	_, err := http.Post(item.WebhookURL, "application/json; charset=utf-8", data)
	if err != nil {
		fmt.Println("Error invoking webhook")
	}
}

// TestHooks api
func TestHooks(w http.ResponseWriter, r *http.Request) {
	var curr currency.DataList
	err2 := database.C("currency").Find(nil).Sort("_id").One(&curr)

	if err2 != nil {
		ErrorWithJSON(w, "Internal error", http.StatusInternalServerError)
		return
	}

	var results []WebHook
	dbErr := database.C("webhooks").Find(bson.M{}).All(&results)

	if dbErr != nil {
		// TODO: Do something about the error
	} else {
		for i := range results {
			item := results[i]
			curr := curr.As(item.BaseCurrency).To(item.TargetCurrency)
			InvokeWebHook(curr, item)
		}
	}
	fmt.Fprint(w, "Invoking webooks")
}

// GetLatestRates api
func GetLatestRates(w http.ResponseWriter, r *http.Request) {
	var latest LatestRates
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&latest)

	if err != nil {
		ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
		return
	}

	var curr currency.DataList
	err2 := database.C("currency").Find(nil).Sort("-_id").One(&curr)

	if err2 != nil {
		ErrorWithJSON(w, "Internal error", http.StatusInternalServerError)
		return
	}
	rate := curr.As(latest.BaseCurrency).To(latest.TargetCurrency)
	fmt.Fprint(w, rate.Rate)
}

// DeleteWebhook api
func DeleteWebhook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var webhook WebHook
	recordId := bson.ObjectIdHex(vars["id"])
	err := database.C("webhooks").FindId(recordId).One(&webhook)

	if err != nil {
		ErrorWithJSON(w, "Not found", http.StatusNotFound)
	} else {
		err2 := database.C("webhooks").RemoveId(recordId)
		if err2 != nil {
			ErrorWithJSON(w, "Internal server error", http.StatusInternalServerError)
		} else {
			WriteJSONResponse(w, webhook)
		}

	}
}

// GetWebhookData api
func GetWebhookData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var webhook WebHook
	recordId := bson.ObjectIdHex(vars["id"])
	err := database.C("webhooks").FindId(recordId).One(&webhook)

	if err != nil {
		ErrorWithJSON(w, "Not found", http.StatusNotFound)
	} else {
		WriteJSONResponse(w, webhook)
	}
}

// AverageGetExchangeRates api
func GetExchange(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Exchange")
}

// PostExchange api
func PostExchange(w http.ResponseWriter, r *http.Request) {

	var webhook WebHook
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&webhook)
	if err != nil {
		ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
		return
	}

	webhook.ID = bson.NewObjectId()

	database.C("webhooks").Insert(webhook)

	fmt.Fprintf(w, webhook.ID.Hex())
}

// WriteJSONResponse api
func WriteJSONResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	output, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		ErrorWithJSON(w, "Internal server Error", http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, string(output))
	}

}

// ErrorWithJSON api
func ErrorWithJSON(w http.ResponseWriter, message string, code int) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	var errorMsg ErrorCode
	errorMsg.Message = message
	errorMsg.Code = code

	output, err := json.MarshalIndent(errorMsg, "", "    ")

	if err != nil {
		ErrorWithJSON(w, "Internal server Error", http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, string(output))
	}
}
