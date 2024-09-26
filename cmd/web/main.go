package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type logger struct {
	infoLogger *log.Logger
	errLogger  *log.Logger
}

func main() {
	host := flag.String("host", ":5000", "Http Port")
	flag.Parse()

	mux := http.NewServeMux()

	// creating an combined object of both these loggers so then can be injected
	appLogger := &logger{
		// Info log instance
		infoLogger: log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC),

		// Error log instance
		errLogger: log.New(os.Stdout, "ERROR\t", log.Ldate|log.LUTC|log.Lshortfile),
	}

	// Get the files from the directory
	fileServer := http.FileServer(http.Dir("../../ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Home Page Request
	mux.HandleFunc("/", appLogger.home)

	// POST REQ to create a snippet
	mux.HandleFunc("/snippet/create", appLogger.snippetCreate)

	// GET REQ to view a snippet with snippetId
	mux.HandleFunc("/snippet/view", appLogger.snippetView)

	appLogger.infoLogger.Printf("Starting server on port %s", *host)

	// new server struct so even http error will use the err log instance
	server := &http.Server{
		Addr:     *host,
		Handler:  mux,
		ErrorLog: appLogger.errLogger,
	}

	serverStartErr := server.ListenAndServe()
	appLogger.errLogger.Fatal(serverStartErr)
}
