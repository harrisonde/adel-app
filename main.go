package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/harrisonde/adel"
)

type application struct {
	App        *adel.Adel
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	a := initApplication()
	a.App.ListenAndServe()
}
