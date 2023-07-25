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

/*
|--------------------------------------------------------------------------
| Graceful Termination
|--------------------------------------------------------------------------
|
| Here is where the wait group is invoked and all items in that were
| registered ask the application to wait until each task for the is done.
| These tasks will block the application until they are complete. For
| example, the application to wait until we have finished sending mail,
| add the mail to wg (i.e., wait group) and when complete call wg.Done()
|
*/
func (a *application) shutdown() {
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
