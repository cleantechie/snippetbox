package main

import (
	"fmt"
	"net/http"
)

// function to append secureHeaders to every request the server recieves
func secureHeaders(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			fmt.Println("Adding secureHeaders to the request")
			response.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
			response.Header().Set("Referrer-Policy", "origin-when-cross-origin")
			response.Header().Set("X-Content-Type-Options", "nosniff")
			response.Header().Set("X-Frame-Options", "deny")
			response.Header().Set("X-XSS-Protection", "0")
			// any code written above is executed on the way down the chain before handling nextHandler
			// if return is invoked here the request wouldnt be passed to the nextHandler
			nextHandler.ServeHTTP(response, request)
			// any code written below is executed on the way back up after handling the above nextHandler
			fmt.Println("Successfully added secureHeaders")
		})
}

// method that logs required information of the request
func (app *application) logRequest(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			app.infoLogger.Printf("%s - %s %s %s", request.RemoteAddr, request.Proto, request.Method, request.URL)
			nextHandler.ServeHTTP(response, request)
		})
}

// method to handle panic
func (app *application) handlePanic(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(res http.ResponseWriter, req *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					app.errLogger.Printf("Gracefully handling the panic")
					res.Header().Set("Connection", "close")
					app.severError(res, fmt.Errorf("%s", err))
				}
			}()
			nextHandler.ServeHTTP(res, req)
		})
}
