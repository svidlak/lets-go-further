package main

import (
	"fmt"
	"net/http"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := map[string]string{
			"RemoteAddr": r.RemoteAddr,
			"Method":     r.Method,
			"Proto":      r.Proto,
			"Url":        r.URL.RequestURI(),
		}
		app.logger.Info("Incoming request:", params)

		next.ServeHTTP(w, r)
	})
}

func (app *application) middlewaresWrapper(handler http.HandlerFunc) http.Handler {
	return app.logRequest(app.recoverPanic(handler))
}
