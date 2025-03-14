package server

import (
	"net/http"
	"time"
)

func New(mux *http.ServeMux, port string) *http.Server {
	return &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
}
