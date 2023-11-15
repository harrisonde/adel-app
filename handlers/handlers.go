package handlers

import (
	"myapp/data"
	"net/http"

	"github.com/harrisonde/adele-framework"
)

type Handlers struct {
	App    *adele.Adele
	Models data.Models
}

/*
|--------------------------------------------------------------------------
| Handlers
|--------------------------------------------------------------------------
|
| Here is where you can add your handlers for the application. These
| handlers are called from your routes.go files.
|
*/

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
