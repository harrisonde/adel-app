package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/harrisonde/adel"
)

type application struct {
	App        *adel.Adel
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
	wg         sync.WaitGroup
}

func main() {
	a := initApplication()
	go a.listenForShutdown()
	err := a.App.ListenAndServe()
	a.App.ErrorLog.Println(err)
}

func (a *application) shutdown() {
	// put all clean up tasks here
	// ...
	// TODO:
	// For example, we might ask the application to wait until
	// we have finished sending email out. To get that done
	// add something to the wg and when it is finished sending
	// say wg.done();

	// block until the wait group is empty
	a.wg.Wait()
}

func (a *application) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	a.App.InfoLog.Println("Received signal", s.String())
	a.shutdown()

	os.Exit(0)
}
