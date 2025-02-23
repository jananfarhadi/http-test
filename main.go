package main

import (
	"github.com/gorilla/mux"
	"github.com/jananfarhadi/http-test/handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/test", handlers.TestHandler).Methods(http.MethodGet)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
