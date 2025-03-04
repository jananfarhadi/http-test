package main

import (
	"github.com/jananfarhadi/http-test/router"
	"log"
)

func main() {
	r := router.NewRoute()

	log.Println("Starting server on", HostAddr)
	log.Fatal(NewHttpService(r).ListenAndServe())
}
