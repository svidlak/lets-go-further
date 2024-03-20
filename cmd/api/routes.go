package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /healthcheck", app.publicRoutesMiddlewaresWrapper(app.healthcheckHandler))

	mux.Handle("GET /v1/movies", app.privateRoutesMiddlewaresWrapper("movies:read", app.listMoviesHandler))
	mux.Handle("POST /v1/movies", app.privateRoutesMiddlewaresWrapper("movies:write", app.createMovieHandler))
	mux.Handle("GET /v1/movies/{id}", app.privateRoutesMiddlewaresWrapper("movies:read", app.showMovieHandler))
	mux.Handle("PATCH /v1/movies/{id}", app.privateRoutesMiddlewaresWrapper("movies:write", app.updateMovieHandler))
	mux.Handle("DELETE /v1/movies/{id}", app.privateRoutesMiddlewaresWrapper("movies:write", app.deleteMovieHandler))

	mux.Handle("POST /v1/users", app.publicRoutesMiddlewaresWrapper(app.registerUserHandler))
	mux.Handle("PUT /v1/users/activated", app.publicRoutesMiddlewaresWrapper(app.activateUserHandler))

	mux.Handle("POST /v1/tokens/authentication", app.publicRoutesMiddlewaresWrapper(app.createAuthenticationTokenHandler))

	return mux
}
