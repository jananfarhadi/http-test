package handlers

import (
	"encoding/json"
	"net/http"
)

func TestHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/Json")
	rw.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(rw).Encode(map[string]interface{}{
		"data": "API Test",
	})
}

func HandshakeHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(rw).Encode(map[string]string{
			"message": "Handshake successful! 🤝",
		})
	} else {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(rw).Encode(map[string]string{
			"error": "Method not allowed",
		})
	}
}
