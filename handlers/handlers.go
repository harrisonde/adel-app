package handlers

import (
	"myapp/data"
	"net/http"

	"github.com/harrisonde/adel"
)

type Handlers struct {
	App    *adel.Adel
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
