package types

import (
	"bytes"
	"database/sql"
	"html/template"

	"github.com/gorilla/sessions"
)

type App struct {
	Session   *sessions.CookieStore
	DB        *sql.DB
	Templates map[string]string
}

func (a *App) GetHtml(name string, data map[string]string) string {
	var w bytes.Buffer
	tmpl, err := template.ParseFiles("web/templates/" + name + ".html")

	if err != nil {
		panic(err)
	}

	tmpl.Execute(&w, data)
	return w.String()
}
