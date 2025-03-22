package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jananfarhadi/http-test/helper"
	"net/http"
	"strings"
)

type (
	UserRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Age       int    `json:"age"`
	}

	UserResponse struct {
		Username   string `json:"username"`
		NameFamily string `json:"name_family"`
		Age        int    `json:"age"`
	}
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.HttpResponse(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "Method not allowed",
		})
		return
	}

	var user UserRequest

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		helper.HttpResponse(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	defer r.Body.Close()

	response := UserResponse{
		Username:   fmt.Sprintf("%s_%s", strings.ToLower(user.FirstName), strings.ToLower(user.LastName)),
		NameFamily: user.FirstName + " " + user.LastName,
		Age:        user.Age,
	}

	helper.HttpResponse(w, http.StatusCreated, response)
}
