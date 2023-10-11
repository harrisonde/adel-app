package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	/*
		|--------------------------------------------------------------------------
		| Global Middleware
		|--------------------------------------------------------------------------
		|
		| Here is where you can add your global Middleware for the application.
		| These middleware are called on each request.
		|
	*/
	a.App.Routes.Use(a.Middleware.CheckForMaintenanceMode)

	/*
		|--------------------------------------------------------------------------
		| Static Routes
		|--------------------------------------------------------------------------
		|
		| Here is where you can add your static routes for the application.
		|
	*/
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))
	a.App.Routes.Mount("/", a.WebRoutes())
	a.App.Routes.Mount("/api", a.ApiRoutes())

	return a.App.Routes
}
