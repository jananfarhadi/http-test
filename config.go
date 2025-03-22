package main

import (
	"net/http"
	"time"
)

const (
	HostAddr     = "0.0.0.0:8080"
	WriteTimeout = 15 * time.Second
	ReadTimeout  = 15 * time.Second
)

func NewHttpService(handler http.Handler) *http.Server {
	return &http.Server{
		Handler:      handler,
		Addr:         HostAddr,
		WriteTimeout: WriteTimeout,
		ReadTimeout:  ReadTimeout,
	}
}
