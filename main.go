package main

import (
	"errors"
	"myapp/cmd"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"net"
	"net/rpc"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/harrisonde/adele-framework"
)

var maintenanceMode bool

type application struct {
	App        *adele.Adele
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
	Commands   *cmd.Commands
	wg         sync.WaitGroup
}

func main() {

	a := initApplication()
	go a.listenForShutdown()
	go a.listenRPC()
	a.jobsSchedule()
	err := a.App.ListenAndServe()
	a.App.ErrorLog.Println(err)
}

/*
|--------------------------------------------------------------------------
| Scheduler
|--------------------------------------------------------------------------
|
| Here is where you may add jobs to the scheduler. Any jobs added will be
| called by the scheduler using the defined interval. You may use one of
| several pre-defined schedules in place of a cron expression (i.e., @yearly,
| @monthly, @weekly, @daily, @hourly and @every <duration>).
|
*/
func (a *application) jobsSchedule() {

	_, err := a.App.Scheduler.AddFunc("* * * * *", func() { a.Commands.RefreshGitHubToken() })

	if err != nil {
		a.App.ErrorLog.Println(err)
	}
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
	a.App.InfoLog.Println("received signal", s.String())
	a.shutdown()

	os.Exit(0)
}

type RPCServer struct {
	App application
}

func (r *RPCServer) MaintenanceMode(inMaintenanceMode bool, resp *string) error {
	res := r.App.setMaintenanceMode(inMaintenanceMode)
	*resp = "server is " + res

	return nil
}

func (a *application) setMaintenanceMode(inMaintenanceMode bool) string {
	if inMaintenanceMode {
		a.App.MaintenanceMode = true
		return "down"
	} else {
		a.App.MaintenanceMode = false
		return "up"
	}
}

func (r *RPCServer) Command(args []string, resp *string) error {

	if len(args) != 4 {
		return errors.New("invalid arguments passed to rpc server in command")
	}

	arg1 := args[0]
	arg2 := args[1]
	arg3 := args[2]
	options := strings.Split(args[3], ",")

	*resp = r.App.Commands.Execute(arg1, arg2, arg3, options)
	return nil
}

func (a *application) listenRPC() {

	a.App.InfoLog.Println("os.Getenv RPC_PORT:", os.Getenv("RPC_PORT"))

	if os.Getenv("RPC_PORT") != "" {
		a.App.InfoLog.Println("Starting RPC server on port", os.Getenv("RPC_PORT"))

		s := new(RPCServer)

		// Provide access to the Adele package
		s.App = *a

		err := rpc.Register(s)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		listen, err := net.Listen("tcp", "127.0.0.1:"+os.Getenv("RPC_PORT"))
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}
		for {
			rpcConn, err := listen.Accept()
			if err != nil {
				continue
			}
			go rpc.ServeConn(rpcConn)

		}
	}
}
