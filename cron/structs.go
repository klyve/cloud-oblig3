package cron

import "gopkg.in/mgo.v2/bson"

// Currency contains the currencies
type Currency struct {
	ID    bson.ObjectId      `json:"id" bson:"_id"`
	Base  string             `json:"base" bson:"base"`
	Date  string             `json:"date" bson:"date"`
	Rates map[string]float32 `json:"rates" bson:"rates"`
}
