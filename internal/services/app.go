package services

import (
	"database/sql"
	"os"

	"github.com/gorilla/sessions"
	"github.com/twoteesgh/go-library/internal/types"
)

func NewAppService(db *sql.DB) *types.App {
	key := []byte(os.Getenv("APP_KEY"))
	store := sessions.NewCookieStore(key)
	templates := map[string]string{
		"guest": "web/templates/partials/guest.html",
		"auth":  "web/templates/partials/auth.html",
	}

	return &types.App{
		Session:   store,
		DB:        db,
		Templates: templates,
	}
}
