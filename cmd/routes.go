package cmd

import (
	"net/http"
	"reflect"
	"runtime"

	"github.com/fatih/color"
	"github.com/go-chi/chi/v5"
	"github.com/harrisonde/adele-framework"
	"github.com/rodaine/table"
)

var RoutesCommand = &adele.Command{
	Name: "route",
	Help: "list all routes for the application",
}

func (c *Commands) List() string {

	type Middleware struct {
		FunctionName string
		Name         string
	}

	type Route struct {
		Method          string
		Path            string
		Middleware      map[int]Middleware
		MiddlewareCount int
	}

	type routes map[int]Route

	type middlewaresFn map[int]Middleware

	routeData := routes{}
	pointer := 1

	chi.Walk(c.App.Routes, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {

		// ... middleware
		middlewareData := middlewaresFn{}
		for index, handlerFunc := range middlewares {
			name := runtime.FuncForPC(reflect.ValueOf(handlerFunc).Pointer()).Name()
			//color.Yellow(fmt.Sprintf("middleware func name %s", name))
			middlewareData[index] = Middleware{
				FunctionName: name,
				Name:         name,
			}
		}

		// build up the route data
		routeData[pointer] = Route{
			Method:          method,
			Path:            route,
			MiddlewareCount: len(middlewares),
			Middleware:      middlewareData,
		}
		pointer++
		return nil
	})

	// Build the table
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Domain", "Method", "URI", "Middleware")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for pointer, route := range routeData {
		if pointer == 1 {
			tbl.AddRow(c.App.Server.URL, route.Method, route.Path, route.MiddlewareCount)
		} else {
			tbl.AddRow("", route.Method, route.Path, route.MiddlewareCount)
		}
	}

	// Print table to the cli
	tbl.Print()

	return ""
}
