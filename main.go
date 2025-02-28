package main

import (
	"github.com/gorilla/mux"
	"log"
)

func main() {
	r := mux.NewRouter()
	log.Fatal(NewHttpService(r).ListenAndServe())
}
