package services

import (
	"database/sql"
	"os"

	"github.com/gorilla/sessions"
)

type App struct {
	Session   *sessions.CookieStore
	DB        *sql.DB
	Templates map[string]string
}

func NewAppService(db *sql.DB) *App {
	key := []byte(os.Getenv("APP_KEY"))
	store := sessions.NewCookieStore(key)

	templates := map[string]string{
		"guest": "views/templates/guest.html",
		"auth":  "views/templates/auth.html",
	}

	return &App{
		Session:   store,
		DB:        db,
		Templates: templates,
	}
}
