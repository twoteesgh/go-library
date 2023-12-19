package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/twoteesgh/go-library/services"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	app *services.App
}

type User struct {
	id         int
	name       string
	email      string
	password   string
	created_at string
	updated_at string
}

func NewAuthHandler(app *services.App) *AuthHandler {
	return &AuthHandler{
		app: app,
	}
}

func (h *AuthHandler) ShowRegisterPage(w http.ResponseWriter, r *http.Request) {
	if tmpl, err := template.ParseFiles(
		"views/register.html",
		h.app.Templates["guest"],
	); err != nil {
		panic(err)
	} else {
		tmpl.Execute(w, nil)
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	user := &User{
		name:     r.FormValue("name"),
		email:    r.FormValue("email"),
		password: hashPassword(r.FormValue("password")),
	}

	_, err := h.app.DB.Exec(`
		INSERT INTO users (name, email, password)
		VALUES (?, ?, ?)
	`, user.name, user.email, user.password)

	if err != nil {
		fmt.Fprint(w, `<p class="text-red-700 font-semibold">Fail</p>`)
	} else {
		fmt.Fprint(w, `<p class="text-green-700 font-semibold">Success</p>`)
	}
}

func (h *AuthHandler) ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	if tmpl, err := template.ParseFiles(
		"views/login.html",
		h.app.Templates["guest"],
	); err != nil {
		panic(err)
	} else {
		tmpl.Execute(w, nil)
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	//
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	//
}

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
