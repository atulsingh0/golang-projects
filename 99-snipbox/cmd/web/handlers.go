package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// Using below middileware function for handling the /static uri
// for blocking URL injections
func neuter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// URL should not have suffix "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := ts.ExecuteTemplate(w, "base", nil); err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// w.Write([]byte("Hello from home page!"))
}

func (app *application) snipView(w http.ResponseWriter, r *http.Request) {

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

func (app *application) snipCreate(w http.ResponseWriter, r *http.Request) {
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
