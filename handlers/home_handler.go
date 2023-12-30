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
	session, _ := h.app.Session.Get(r, "auth")
	wrapper := h.app.Templates["guest"]

	if session.Values["user"] != nil {
		wrapper = h.app.Templates["auth"]
	}

	if tmpl, err := template.ParseFiles(
		"views/home.html",
		wrapper,
	); err != nil {
		panic(err)
	} else {
		err := tmpl.Execute(w, session.Values["user"])
		if err != nil {
			panic(err)
		}
	}
}
