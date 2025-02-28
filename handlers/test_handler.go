package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func TestHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "API is working! ")
	rw.Header().Set("Content-Type", "application/Json")
	rw.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(rw).Encode(map[string]interface{}{"data": "API Test"})

	return
}
