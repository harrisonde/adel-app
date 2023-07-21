package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) ApiRoutes() http.Handler {

	// Create sub-router to be mounted
	r := chi.NewRouter()

	r.Route("/api", func(mux chi.Router) {

		// Add API Routes here
		// ...

	})

	return r
}
