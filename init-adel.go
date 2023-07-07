package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"

	"github.com/harrisonde/adel"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	a := &adel.Adel{} // init Adel with reference to Adel
	err = a.New(path)
	if err != nil {
		log.Fatal(err)
	}

	a.AppName = "myapp"

	myMiddleware := &middleware.Middleware{
		App: a,
	}

	myHandlers := &handlers.Handlers{
		App: a,
	}

	app := &application{
		App:        a,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	// Add our application routes to the default routes
	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)

	myHandlers.Models = app.Models

	app.Middleware.Models = app.Models

	return app
}
