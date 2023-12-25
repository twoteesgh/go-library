package handlers

import (
	"html/template"
	"net/http"

	"github.com/twoteesgh/go-library/services"
)

type HomeHandler struct {
	app *services.App
}

func NewHomeHandler(app *services.App) *HomeHandler {
	return &HomeHandler{
		app: app,
	}
}

func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	if tmpl, err := template.ParseFiles(
		"views/home.html",
		h.app.Templates["guest"],
	); err != nil {
		panic(err)
	} else {
		session, _ := h.app.Session.Get(r, "auth")
		err := tmpl.Execute(w, session.Values["user"].(User))
		if err != nil {
			panic(err)
		}
	}
}
