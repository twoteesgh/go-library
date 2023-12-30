package handlers

import (
	"net/http"

	"github.com/twoteesgh/go-library/services"
)

type UserHandler struct {
	app *services.App
}

func NewUserHandler(app *services.App) *UserHandler {
	return &UserHandler{
		app: app,
	}
}

func (h *UserHandler) Borrow(w http.ResponseWriter, r *http.Request) {
	//
}
