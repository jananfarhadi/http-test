package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jananfarhadi/http-test/helper"
	"net/http"
)

func GetUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	msg := fmt.Sprintf("User Info: ID: %s", vars["id"])

	helper.HttpResponse(rw, http.StatusOK, map[string]interface{}{
		"detail": msg,
	})

	return
}

func GetUserInfo(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		helper.HttpResponse(rw, http.StatusUnprocessableEntity, map[string]interface{}{
			"error": "name is empty",
		})

		return
	}

	age := r.URL.Query().Get("age")
	if age == "" {
		helper.HttpResponse(rw, http.StatusUnprocessableEntity, map[string]interface{}{
			"error": "age is empty",
		})

		return
	}

	helper.HttpResponse(rw, http.StatusOK, map[string]interface{}{
		"name": name,
		"age":  age,
	})
	return
}
