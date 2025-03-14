package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parsing form error", err)
		return
	}
	fmt.Fprintf(w, "Post request is successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name is %s \n", name)
	fmt.Fprintf(w, "Address is %s \n", address)

}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting Server at port 9000")

	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}

	// http.HandleFunc("/", greet)
	// http.ListenAndServe(":8080", nil)
}
