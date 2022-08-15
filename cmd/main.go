package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"parser_test/internal/router"
	"time"
)

func main() {
	r := mux.NewRouter()
	router.InitRouter(r)
	runMuxServer(r)
}

func runMuxServer(router *mux.Router) {
	startServer := "0.0.0.0:8089"
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
