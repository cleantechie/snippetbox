package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.harshasv.net/internal/datamodels"
)

// handler method to display latest snippets on home page
func (app *application) home(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(response, request)
		return
	}
	latestSnippets, err := app.snippetModel.GetLatestiSnippets()
	if err != nil {
		app.severError(response, err)
	}

	app.renderTemplate(response, http.StatusOK, "home.html", &templateData{
		LatestSnippets: latestSnippets,
	})
}

// handler method to create a new snippet
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

// handler method to handle displaying of specific snippet
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
	app.renderTemplate(response, http.StatusOK, "view.html", &templateData{
		Snippet: snippet,
	})

	app.infoLogger.Printf("Displaying a specfic snippet with Id %d....", snippet.ID)
	fmt.Fprintf(response, "%+v", snippet)
}
