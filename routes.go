package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// Middleware must come before any routes

	// Static Routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	// Mount Web and API routes
	a.App.Routes.Mount("/", a.WebRoutes())
	a.App.Routes.Mount("/api", a.ApiRoutes())

	return a.App.Routes
}
