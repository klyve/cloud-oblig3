package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"github.com/klyve/cloud-oblig3/api"
	"github.com/klyve/cloud-oblig3/bot"
	"github.com/klyve/cloud-oblig3/cron"
	"gopkg.in/mgo.v2"
)

// Config struct for db and port config
type Config struct {
	Port            int    `env:"PORT" envDefault:"3000"`
	MongoDBHost     string `env:"MONGODB_HOST" envDefault:"localhost"`
	MongoDBUsername string `env:"MONGODB_USERNAME" envDefault:""`
	MongoDBPassword string `env:"MONGODB_PASSWORD" envDefault:""`
	MongoDBPort     int    `env:"MONGODB_PORT" envDefault:"27017"`
	MongoDBDatabase string `env:"MONGODB_DATABASE" envDefault:"NWA"`
}

// type config struct {
// 	Home         string        `env:"HOME"`
// 	Port         int           `env:"PORT" envDefault:"3000"`
// 	IsProduction bool          `env:"PRODUCTION"`
// 	Hosts        []string      `env:"HOSTS" envSeparator:":"`
// 	Duration     time.Duration `env:"DURATION"`
// }

// GetHomePage Endpoint
func GetHomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func createMongoDBURI(cfg Config) string {

	p := strconv.Itoa(cfg.MongoDBPort)
	portAddr := ":" + p
	if cfg.MongoDBUsername == "" {
		return "mongodb://" + cfg.MongoDBHost + portAddr
	}
	var ret string
	ret += "mongodb://"
	ret += cfg.MongoDBUsername
	ret += ":" + cfg.MongoDBPassword
	ret += "@" + cfg.MongoDBHost
	ret += portAddr

	return ret
}

// Init the app
func Init(prod bool) {
	// cron.Init()
	router := mux.NewRouter()

	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	p := strconv.Itoa(cfg.Port)
	portAddr := ":" + p

	fmt.Println(cfg.MongoDBHost)
	fmt.Println(cfg.Port)

	router.HandleFunc("/", GetHomePage).Methods("GET")
	mongoURI := createMongoDBURI(cfg)
	fmt.Println(mongoURI)
	session, err := mgo.Dial(mongoURI)
	if err != nil {
		log.Fatal("Could not connect to the database")
	}

	database := session.DB(cfg.MongoDBDatabase)
	cron.Init(database)
	api.Init(router, database)
	bot.Init(router, database)
	if prod == true {
		log.Fatal(http.ListenAndServe(portAddr, router))
		fmt.Printf("Connected on port %s", p)
	}
}

// Main entrypoint
func main() {
	Init(true)
}
