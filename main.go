package main

import (
	"github.com/gorilla/mux"
	"github.com/jananfarhadi/http-test/handlers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/test", handlers.TestHandler).Methods(http.MethodGet)

	log.Fatal(NewHttpService(r).ListenAndServe())
}
