package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) ApiRoutes() http.Handler {

	r := chi.NewRouter()

	r.Route("/api", func(mux chi.Router) {

		/*
		|--------------------------------------------------------------------------
		| API Routes
		|--------------------------------------------------------------------------
		|
		| Here is where you can add your API routes for the application. These
		| routes are loaded by the router.
		|
		*/

	})

	return r
}
