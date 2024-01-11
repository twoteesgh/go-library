package handlers

import (
	"net/http"

	"github.com/twoteesgh/go-library/internal/types"
)

type UserHandler struct {
	app *types.App
}

func NewUserHandler(app *types.App) *UserHandler {
	return &UserHandler{
		app: app,
	}
}

func (h *UserHandler) Borrow(w http.ResponseWriter, r *http.Request) {
	//
}
