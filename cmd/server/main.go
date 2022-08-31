package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"parser_test/config"
	"parser_test/internal/bl"
	"parser_test/internal/router"
	"parser_test/internal/store/mongo"
	"time"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatal(err)
	}

	mongoClient, err := mongo.NewMongoClient()
	if err != nil {
		log.Fatal("[ERROR] : ", "MongoClient error", err.Error())
	}
	mongoRepo, err := mongo.NewMongoRepo(mongoClient)
	if err != nil {
		log.Fatal("[ERROR] : ", "MongoRepo error", err.Error())
	}
	serverBl := bl.NewBL(mongoRepo)

	r := mux.NewRouter()
	router.InitRouter(r, serverBl)
	runMuxServer(r)
}

func runMuxServer(router *mux.Router) {
	startServer := config.GetSettings().Host + ":" + config.GetSettings().Port
	srv := &http.Server{
		Handler: router,
		Addr:    startServer,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	log.Println("[INFO] : ", "Server started ", startServer)
	if err := srv.ListenAndServe(); err != nil {
		log.Println("[ERROR] : ", "ListenAndServe error ", err.Error())
	}
}
