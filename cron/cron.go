package cron

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/klyve/cloud-oblig2/api"
	"github.com/klyve/cloud-oblig2/currency"

	"github.com/robfig/cron"
	"gopkg.in/mgo.v2"
)

var collection *mgo.Collection
var database *mgo.Database

// Init cron jobs
func Init(db *mgo.Database) {
	collection = db.C("currency")
	database = db

	c := cron.New()
	CronJob()
	c.AddFunc("@every 24h", CronJob)
	c.Start()
}

// CronJob item
func CronJob() {
	latest := currency.FetchLatest()
	data := currency.Currency{
		ID:    bson.NewObjectId(),
		Base:  latest.Base,
		Date:  latest.Date,
		Rates: latest.Rates,
	}
	err := collection.Insert(&data)
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	var results []api.WebHook
	dbErr := database.C("webhooks").Find(bson.M{}).All(&results)
	if dbErr != nil {
		// TODO: Do something about the error
	} else {
		for i := range results {
			item := results[i]
			curr := latest.As(item.BaseCurrency).To(item.TargetCurrency)
			if curr.Rate > item.MaxTriggerValue || curr.Rate < item.MinTriggerValue {
				api.InvokeWebHook(curr, item)
			} else {
				fmt.Println("Results All: ", item.BaseCurrency)
			}
		}
	}
	fmt.Println("Cron job")
}
