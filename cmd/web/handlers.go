package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(response, request)
		return
	}
	TEMPLATE, ERROR := template.ParseFiles(
		"../../ui/html/base.html",
		"../../ui/html/partials/nav.html",
		"../../ui/html/pages/home.html",
	)
	if ERROR != nil {
		log.Print(ERROR.Error())
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
	}
	ERROR = TEMPLATE.ExecuteTemplate(response, "base", nil)
	if ERROR != nil {
		log.Print(ERROR.Error())
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
	}
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
	snippetId, ERROR := strconv.Atoi(request.URL.Query().Get("snippetId"))
	if ERROR != nil || snippetId < 1 {
		http.NotFound(response, request)
		return
	}
	fmt.Fprintf(response, "Displaying a specfic snippet with Id %d....", snippetId)
	response.Write([]byte("Display a new Snippet"))
}
