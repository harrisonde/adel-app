package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) WebRoutes() http.Handler {

	r := chi.NewRouter()

	r.Route("/", func(mux chi.Router) {

		// Add Web Routes here
		// ...

		r.Get("/", a.Handlers.Home)

	})
	return r
}
