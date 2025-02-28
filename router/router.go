package router

import (
	"github.com/gorilla/mux"
	"github.com/jananfarhadi/http-test/handlers"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/test", handlers.TestHandler).Methods(http.MethodGet)

	return r
}
