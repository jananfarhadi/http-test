package handlers

import (
	"github.com/jananfarhadi/http-test/helper"
	"net/http"
)

func HandshakeHandler(rw http.ResponseWriter, r *http.Request) {
	lang := r.Header.Get("lang")
	if r.Method == http.MethodPost {
		helper.HttpResponse(rw, http.StatusOK, map[string]string{
			"lang":    lang,
			"message": "Handshake successful! 🤝",
		})

		return
	}

	helper.HttpResponse(rw, http.StatusMethodNotAllowed, map[string]string{
		"error": "Method not allowed",
	})
}
