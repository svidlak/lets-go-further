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

	mux.Handle("POST /v1/users", app.middlewaresWrapper(app.registerUserHandler))
	mux.Handle("PUT /v1/users/activated", app.middlewaresWrapper(app.activateUserHandler))

	mux.Handle("POST /v1/tokens/authentication", app.middlewaresWrapper(app.createAuthenticationTokenHandler))

	return mux
}
