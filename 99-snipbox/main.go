package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snip/view", snipView)
	mux.HandleFunc("/snip/create", snipCreate)

	log.Println("Starting server on :4000")

	if err := http.ListenAndServe(`:4000`, mux); err != nil {
		log.Fatal("Unable to start server", err.Error())
	}
}

func home(w http.ResponseWriter, r *http.Request) {

	// URL should be strictly "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from home page!"))
}

func snipView(w http.ResponseWriter, r *http.Request) {

	// Only GET method is valid
	if r.Method != http.MethodGet {
		// Use http.Error method
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method Not Allowed"))
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Viewing a snippet"))
}

func snipCreate(w http.ResponseWriter, r *http.Request) {
	// Only POST method is valid
	if r.Method != http.MethodPost {
		// All the changes in header should be done before calling WriteHeader method
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method Not Allowed"))
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Creating a snippet"))
}
