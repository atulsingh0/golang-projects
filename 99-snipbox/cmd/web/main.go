package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	// Define Arguements
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new instance of our application struct, containing the
	// dependencies.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Define Server
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", neuter(fileServer)))

	// Register Handlers
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snipView)
	mux.HandleFunc("/snippet/create", app.snipCreate)

	log.Printf("Starting server on %s\n", *addr)

	// Starting Server
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal("Unable to start server", err.Error())
	}
}
