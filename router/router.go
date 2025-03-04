package router

import (
	"github.com/gorilla/mux"
	"github.com/jananfarhadi/http-test/handlers"
	"net/http"
)

func NewRoute() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/test", handlers.TestHandler).Methods(http.MethodGet)
	api.HandleFunc("/handshake", handlers.HandshakeHandler).Methods(http.MethodPost)

	return r
}
