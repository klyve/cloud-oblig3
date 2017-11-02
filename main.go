package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/klyve/cloud-oblig2/api"
	"github.com/klyve/cloud-oblig2/cron"
	"gopkg.in/mgo.v2"
)

type Config struct {
	Port int
}

func GetHomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func Init() {
	// cron.Init()
	router := mux.NewRouter()

	// currency.PrintTo()
	// Config
	cfg := Config{Port: 3000}

	p := strconv.Itoa(cfg.Port)
	portAddr := ":" + p

	router.HandleFunc("/", GetHomePage).Methods("GET")

	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal("Could not connect to the database")
	}

	database := session.DB("oblig2")
	cron.Init(database)
	api.Init(router, database)

	log.Fatal(http.ListenAndServe(portAddr, router))
	fmt.Printf("Connected on port %s", p)
}

func main() {
	Init()
}
