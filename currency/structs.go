package currency

import "gopkg.in/mgo.v2/bson"

// DataList Holds the currency data
type DataList struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float32 `json:"rates"`
}

// Convertion Holds a single from to currency value
type Convertion struct {
	From      string  `json:"from"`
	FromValue float32 `json:"from_value"`
	To        string  `json:"to"`
	ToValue   float32 `json:"to_value"`
	Rate      float32 `json:"rate"`
}

// Currency holds a currency type
type Currency struct {
	ID    bson.ObjectId      `json:"id" bson:"_id"`
	Base  string             `json:"base" bson:"base"`
	Date  string             `json:"date" bson:"date"`
	Rates map[string]float32 `json:"rates" bson:"rates"`
}
