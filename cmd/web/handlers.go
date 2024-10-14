package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"snippetbox.harshasv.net/internal/datamodels"
)

func (app *application) home(response http.ResponseWriter, request *http.Request) {
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
		app.notFound(response)
		return
	}
	snippet, err := app.snippetModel.GetSnippetById(snippetId)
	if err != nil {
		if errors.Is(err, datamodels.ErrNoRecord) {
			app.notFound(response)
		} else {
			app.severError(response, err)
		}
		return
	}
	files := []string{
		"../../ui/html/base.html",
		"../../ui/html/partials/nav.html",
		"../../ui/html/pages/view.html",
	}
	template, templateErr := template.ParseFiles(files...)
	if templateErr != nil {
		app.severError(response, templateErr)
		return
	}

	templateErr = template.ExecuteTemplate(response, "base", snippet)
	if templateErr != nil {
		app.severError(response, templateErr)
		return
	}

	app.infoLogger.Printf("Displaying a specfic snippet with Id %d....", snippetId)

	fmt.Fprintf(response, "%+v", snippet)
	// jsonResult, marshallingErr := json.Marshal(snippet)
	// if marshallingErr != nil {
	// 	app.clientError(response, http.StatusInternalServerError)
	// }
	// response.Header().Set("Content-Type", "application/json")
	// response.Write(jsonResult)
}
