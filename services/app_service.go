package services

import (
	"database/sql"
	"html/template"
	"net/http"
)

type App struct {
	DB        *sql.DB
	Templates map[string]string
}

func NewAppService(db *sql.DB) *App {
	templates := map[string]string{
		"guest": "views/templates/guest.html",
		"auth":  "views/templates/auth.html",
	}

	return &App{
		DB:        db,
		Templates: templates,
	}
}

func (h *App) Home(w http.ResponseWriter, r *http.Request) {
	if tmpl, err := template.ParseFiles(
		"views/home.html",
		h.Templates["guest"],
	); err != nil {
		panic(err)
	} else {
		tmpl.Execute(w, nil)
	}
}
