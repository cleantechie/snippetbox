package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Generic 500 InternalServerError
func (app *application) severError(response http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s \n %s", err.Error(), debug.Stack())
	// logger.Print is a wrapper of Output with by default calldepth 2
	// logger.Output we need to format the trace String
	app.errLogger.Output(2, trace)

	http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Generic 400 BadRequest
func (app *application) clientError(response http.ResponseWriter, statusCode int) {
	http.Error(response, http.StatusText(statusCode), statusCode)
}

// Generic 404 notFound
func (app *application) notFound(response http.ResponseWriter) {
	app.clientError(response, http.StatusNotFound)
}
