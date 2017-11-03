package api

import "gopkg.in/mgo.v2/bson"

// WebHook data type
type WebHook struct {
	ID              bson.ObjectId `json:"id" bson:"_id"`
	WebhookURL      string        `json:"webhookURL"`
	BaseCurrency    string        `json:"baseCurrency"`
	TargetCurrency  string        `json:"targetCurrency"`
	MinTriggerValue float32       `json:"minTriggerValue"`
	MaxTriggerValue float32       `json:"maxTriggerValue"`
}

// WebHookTest data type
type WebHookTest struct {
	ID              string  `json:"id" bson:"_id"`
	WebhookURL      string  `json:"webhookURL"`
	BaseCurrency    string  `json:"baseCurrency"`
	TargetCurrency  string  `json:"targetCurrency"`
	MinTriggerValue float32 `json:"minTriggerValue"`
	MaxTriggerValue float32 `json:"maxTriggerValue"`
}

// WebHookInvoked data type
type WebHookInvoked struct {
	BaseCurrency    string  `json:"baseCurrency"`
	TargetCurrency  string  `json:"targetCurrency"`
	CurrentRate     float32 `json:"currentRate"`
	MinTriggerValue float32 `json:"minTriggerValue"`
	MaxTriggerValue float32 `json:"maxTriggerValue"`
}

// LatestRates data type
type LatestRates struct {
	BaseCurrency   string `json:"baseCurrency"`
	TargetCurrency string `json:"targetCurrency"`
}

// CurrencyResponse data type
type CurrencyResponse struct {
	BaseCurrency   string  `json:"baseCurrency"`
	TargetCurrency string  `json:"targetCurrency"`
	CurrentRate    float32 `json:"currentRate"`
}

// DataListTest data type
type DataListTest struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float32 `json:"rates"`
}
