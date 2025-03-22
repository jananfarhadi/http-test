package helper

import (
	"encoding/json"
	"net/http"
	"time"
)

type response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Timestamp  string      `json:"timestamp"`
	Data       interface{} `json:"data"`
}

func HttpResponse(w http.ResponseWriter, status int, data interface{}) {
	var message string

	if 300 <= status {
		message = "Request process failed"
	} else {
		message = "Request process successfully"
	}

	resp := response{
		StatusCode: status,
		Message:    message,
		Timestamp:  time.Now().Format(time.RFC3339),
		Data:       data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)

	return
}
