package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello from Snippet box"))
}

func snippetCreate(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		response.Header().Set("Allow", "POST")
		http.Error(response, "This Method isnt allowed", http.StatusMethodNotAllowed)

	}
	response.Header().Set("Content-Type", "application/json")
	response.Write([]byte(`{"name" : "test"}`))
}

func snippetView(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Display a new Snippet"))
}

func main() {
	fmt.Println("ertyul")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /home/", home)
	mux.HandleFunc("POST /snippet/create", snippetCreate)
	mux.HandleFunc("GET /snippet/view", snippetView)
	log.Print("Starting server on port 5000")
	serverStartErr := http.ListenAndServe(":5000", mux)
	log.Fatal(serverStartErr)
}
