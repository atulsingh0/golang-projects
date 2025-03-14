package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		log.Println("Hello world!")
		log.Println(r.Header)
		log.Println(r.Body)
		rw.Write([]byte("Hi there"))

	})

	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("Good Bye world!")
	})

	http.ListenAndServe(":9090", nil)
}
