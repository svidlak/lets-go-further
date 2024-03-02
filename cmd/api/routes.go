package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /healthcheck", app.middlewaresWrapper(app.healthcheckHandler))

	mux.Handle("GET /v1/movies", app.middlewaresWrapper(app.listMoviesHandler))
	mux.Handle("POST /v1/movies", app.middlewaresWrapper(app.createMovieHandler))
	mux.Handle("GET /v1/movies/{id}", app.middlewaresWrapper(app.showMovieHandler))
	mux.Handle("PATCH /v1/movies/{id}", app.middlewaresWrapper(app.updateMovieHandler))
	mux.Handle("DELETE /v1/movies/{id}", app.middlewaresWrapper(app.deleteMovieHandler))

	return mux
}
