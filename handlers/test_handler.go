package handlers

import (
	"github.com/jananfarhadi/http-test/helper"
	"net/http"
)

func TestHandler(rw http.ResponseWriter, r *http.Request) {
	helper.HttpResponse(rw, http.StatusOK, map[string]interface{}{
		"detail": "API Test",
	})
}
