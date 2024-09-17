package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("POST /snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/view", snippetView)
	log.Print("Starting server on port 5000")
	serverStartErr := http.ListenAndServe(":5000", mux)
	log.Fatal(serverStartErr)
}
