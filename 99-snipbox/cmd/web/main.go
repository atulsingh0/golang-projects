package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	// Define Arguements
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Define Server
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", neuter(fileServer)))

	// Register Handlers
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snipView)
	mux.HandleFunc("/snippet/create", snipCreate)

	log.Printf("Starting server on %s\n", *addr)

	// Starting Server
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal("Unable to start server", err.Error())
	}
}
