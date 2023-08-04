package main

import (
	"fmt"
	"net/http"
	"strconv"
)

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
		w.Header().Set("Allow", http.MethodGet)

		// Use http.Error method
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("Method Not Allowed"))
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Processing URL Query
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Viewing a snippet"))
	// using Fprintf func
	fmt.Fprintf(w, "Viewing a snippet with ID %d", id)
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
