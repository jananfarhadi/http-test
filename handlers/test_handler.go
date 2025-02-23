package handlers

import (
	"encoding/json"
	"net/http"
)

func TestHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/Json")
	rw.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(rw).Encode(map[string]interface{}{"data": "API Test"})

	return
}
