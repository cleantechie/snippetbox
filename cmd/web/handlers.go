package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func (applogger *logger) home(response http.ResponseWriter, request *http.Request) {
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
		applogger.errLogger.Print(ERROR.Error())
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
	}
	ERROR = TEMPLATE.ExecuteTemplate(response, "base", nil)
	if ERROR != nil {
		applogger.errLogger.Print(ERROR.Error())
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (applogger *logger) snippetCreate(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		response.Header().Set("Allow", "POST")
		http.Error(response, "This Method isnt allowed", http.StatusMethodNotAllowed)

	}
	applogger.infoLogger.Print("Creating an new Snippet")
	response.Header().Set("Content-Type", "application/json")
	response.Write([]byte(`{"name" : "test"}`))
}

func (applogger *logger) snippetView(response http.ResponseWriter, request *http.Request) {
	snippetId, ERROR := strconv.Atoi(request.URL.Query().Get("snippetId"))
	if ERROR != nil || snippetId < 1 {
		http.NotFound(response, request)
		return
	}
	applogger.infoLogger.Printf("Displaying a specfic snippet with Id %d....", snippetId)
	response.Write([]byte("Display a new Snippet"))
}
