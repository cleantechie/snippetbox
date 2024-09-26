package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	host := flag.String("host", ":5000", "Http Port")
	flag.Parse()
	mux := http.NewServeMux()
	// Info log instance
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC)
	// Err log instance
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.LUTC|log.Lshortfile)

	// Get the files from the directory
	fileServer := http.FileServer(http.Dir("../../ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Home Page Request
	mux.HandleFunc("/", home)
	// POST REQ to create a snippet
	mux.HandleFunc("/snippet/create", snippetCreate)
	// GET REQ to view a snippet with snippetId
	mux.HandleFunc("/snippet/view", snippetView)
	infoLog.Printf("Starting server on port %s", *host)

	// new server struct so even http error will use the err log instance
	server := &http.Server {
		Addr: *host,
		Handler: mux,
		ErrorLog: errLog,
	}

	serverStartErr := server.ListenAndServe()
	errLog.Fatal(serverStartErr)
}
