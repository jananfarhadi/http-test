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
	api.HandleFunc("/user", handlers.UserHandler).Methods(http.MethodPost)
	api.HandleFunc("/user/{id:[1][0-5][02468][13579].*}", handlers.GetUser).Methods(http.MethodGet)
	api.HandleFunc("/user-info", handlers.GetUserInfo).Methods(http.MethodGet)
	return r
}
