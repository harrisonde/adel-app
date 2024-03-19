package handlers

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"git.int.86labs.cloud/harrisonde/adel"
	"git.int.86labs.cloud/harrisonde/adel/mailer"
	"git.int.86labs.cloud/harrisonde/adel/render"
	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
)

var adl adel.Adel
var testSession *scs.SessionManager
var testHandlers Handlers

func TestMain(m *testing.M) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Setup session
	testSession = scs.New()
	testSession.Lifetime = 24 * time.Hour
	testSession.Cookie.Persist = true
	testSession.Cookie.SameSite = http.SameSiteLaxMode
	testSession.Cookie.Secure = false

	// Set jet
	var views = jet.NewSet(
		jet.NewOSFileSystemLoader("../views"),
		jet.InDevelopmentMode(),
	)

	myRenderer := render.Render{
		Renderer: "jet",
		RootPath: "../",
		Port:     "4000",
		JetViews: views,
		Session:  testSession,
	}

	adl = adel.Adel{
		AppName:       "myapp",
		Debug:         true,
		Version:       "1.0.0",
		ErrorLog:      errorLog,
		InfoLog:       infoLog,
		RootPath:      "../",
		Routes:        nil,
		Render:        &myRenderer,
		Session:       testSession,
		DB:            adel.Database{},
		JetViews:      views,
		EncryptionKey: adl.RandomString(32),
		Cache:         nil,
		Scheduler:     nil,
		Mail:          mailer.Mail{},
		Server:        adel.Server{},
	}

	testHandlers.App = &adl

	os.Exit(m.Run())
}

// Routes that need to be created when setting up the testing server
// ...
func getRoutes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(adl.SessionLoad)

	// Add any Routes here
	// ...

	mux.Get("/", testHandlers.Home)

	// Expose static content to the testing File Server
	fileServer := http.FileServer(http.Dir("./../public"))
	mux.Handle("/public/*", http.StripPrefix("/public", fileServer))
	return mux
}

func getCtx(req *http.Request) context.Context {
	// Need this to access information from the session
	// without it you cannot access anything that is in the session!
	ctx, err := testSession.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
