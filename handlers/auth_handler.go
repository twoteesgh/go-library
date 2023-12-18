package handlers

import (
	"net/http"

	"github.com/twoteesgh/go-library/services"
)

type AuthHandler struct {
	app *services.App
}

type User struct {
	id         int
	name       string
	email      string
	created_at string
	updated_at string
}

func NewAuthHandler(app *services.App) *AuthHandler {
	return &AuthHandler{
		app: app,
	}
}

func (h *AuthHandler) ShowRegisterPage(w http.ResponseWriter, r *http.Request) {
	//
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	//
}

func (h *AuthHandler) ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	//
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	//
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	//
}
