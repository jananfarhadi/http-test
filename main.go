package main

import (
	"github.com/jananfarhadi/http-test/router"
	"log"
)

func main() {
	r := router.NewRoute()

	log.Fatal(NewHttpService(r).ListenAndServe())
}
