package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(response, request)
		return
	}
	TEMPLATE, ERROR := template.ParseFiles(
		"../../ui/html/base.html",
		"../../ui/html/partials/nav.htl",
		"../../ui/html/pages/home.html",
	)
	if ERROR != nil {
		app.severError(response, ERROR)
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
	}
	ERROR = TEMPLATE.ExecuteTemplate(response, "base", nil)
	if ERROR != nil {
		app.severError(response, ERROR)
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) snippetCreate(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		response.Header().Set("Allow", "POST")
		app.clientError(response, http.StatusMethodNotAllowed)
		return
	}
	title := "Test"
	content := "This was meant to be a test"
	expriesAt := 3
	app.infoLogger.Print("Creating an new Snippet")
	snippedtId, executionErr := app.snippetModel.InsertASnippet(title, content, expriesAt)
	if executionErr != nil {
		app.severError(response, executionErr)
		return
	}

	http.Redirect(response, request, fmt.Sprintf("/snippet/view?snippetId=%d", snippedtId), http.StatusSeeOther)
}

func (app *application) snippetView(response http.ResponseWriter, request *http.Request) {
	snippetId, ERROR := strconv.Atoi(request.URL.Query().Get("snippetId"))
	if ERROR != nil || snippetId < 1 {
		http.NotFound(response, request)
		return
	}
	app.infoLogger.Printf("Displaying a specfic snippet with Id %d....", snippetId)
	response.Write([]byte("Display a new Snippet"))
}
