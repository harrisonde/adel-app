package handlers

import (
	"myapp/data"
	"net/http"
	"strings"

	"github.com/harrisonde/adel"
)

type Handlers struct {
	App    *adel.Adel
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

func (h *Handlers) Inertia(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	template := strings.Replace(url, "/", "", 1)

	if template == "" {
		template = "index"
	} else {

		ok := h.isAlphaAnd(template, "/")
		if !ok {
			h.App.ErrorStatus(w, http.StatusNotFound)
			return
		}

		ok = h.isAlpha(template)
		if ok {
			template = template + "/index"
		}
	}

	template = strings.ToLower(template)
	err := h.renderInertia(w, r, template)
	if err != nil {
		h.App.ErrorLog.Println("error rendering inertia page:", err)
	}
}
