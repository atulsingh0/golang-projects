package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcom to JWT"))
	})

	mux.HandleFunc("/gen", generate)

	err := http.ListenAndServe(":8888", mux)

	if err != nil {
		log.Fatal("Server is unable to start")
	}
}

func generate(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling Req:", r.RequestURI)
	w.Header().Add("Server", "jwt")
	msg := "this is me calling from generate"
	w.Write([]byte(msg))
}
