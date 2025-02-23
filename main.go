package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(rw).Encode(map[string]interface{}{"data": "API Test"})

		return //   چون تابع هندلر به‌طور پیش‌فرض بعد از اجرا تموم میشه. می‌تونی حذفش کنی.
	})

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
