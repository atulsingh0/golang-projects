package main

import (
	"fmt"
	"net/http"
	"websvc/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", helloHandler)
	mux.HandleFunc("POST /hello", helloPostHandler)
	mux.HandleFunc("GET /emp/{id}", getEmpDetails)

	// handling specific host
	// use -H Host: devs.net
	mux.HandleFunc("GET devs.net/", devHandler)

	// order matters
	stacked := middleware.CreateStack(
		middleware.IsAuth,
		middleware.Logging,
	)

	adminstack := middleware.CreateStack(
		middleware.EnsureAdmin,
		middleware.Logging,
	)

	// versioned handle (using sub-router)
	v1serve := http.NewServeMux()
	v1serve.Handle("/v1/", http.StripPrefix("/v1", stacked(mux)))

	// path strict to admin (sub-router)

	admin := http.NewServeMux()
	mux.Handle("/admin/", http.StripPrefix("/admin", adminstack(admin)))

	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	fmt.Println("Listening on port 3000")
	err := server.ListenAndServe()
	if err != nil {
		return
	}

}

func devHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello! Devs"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		return
	}
	return
}

func helloPostHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	w.Write([]byte("Hello " + name + "!, Welcome to PostHandler!"))
}

func getEmpDetails(w http.ResponseWriter, r *http.Request) {
	data := fmt.Sprintf("Hello %s!", r.PathValue("id"))
	w.Write([]byte(data))
}
