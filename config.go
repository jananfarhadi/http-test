package main

import (
	"net/http"
)

const (
	HostAddr     = "127.0.0.1:8080"
	WriteTimeout = 15
	ReadTimeout  = 15
)

func NewHttpService(handler http.Handler) *http.Server {
	return &http.Server{
		Handler:      handler,
		Addr:         HostAddr,
		WriteTimeout: WriteTimeout,
		ReadTimeout:  ReadTimeout,
	}
}
