package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)
	mux.HandleFunc("GET /v1/movies", app.listMoviesHandler)
	mux.HandleFunc("POST /v1/movies", app.createMovieHandler)
	mux.HandleFunc("GET /v1/movies/{id}", app.showMovieHandler)
	mux.HandleFunc("PUT /v1/movies/{id}", app.editMovieHandler)
	mux.HandleFunc("DELETE /v1/movies/{id}", app.deleteMovieHandler)

	return mux
}
