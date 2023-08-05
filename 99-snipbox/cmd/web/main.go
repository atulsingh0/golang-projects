package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", neuter(fileServer)))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snipView)
	mux.HandleFunc("/snippet/create", snipCreate)

	log.Println("Starting server on :4000")

	if err := http.ListenAndServe(`:4000`, mux); err != nil {
		log.Fatal("Unable to start server", err.Error())
	}
}
