package main

import (
	"log"
	"net/http"
	"os"

	"98-web-app-basic/homepage"

	"98-web-app-basic/server"
)

func main() {

	mux := http.NewServeMux()

	logger := log.New(os.Stdout, "web-app: ", log.LstdFlags|log.Lshortfile)
	h := homepage.NewHandlers(logger)

	mux.HandleFunc("/", h.HomeHandler)

	srv := server.New(mux, ":8080")
	logger.Println("Server is starting...")

	err := srv.ListenAndServe()

	if err != nil {
		logger.Fatal(err.Error())
	}
}
