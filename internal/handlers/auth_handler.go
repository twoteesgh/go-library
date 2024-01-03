package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"net/http"

	"github.com/twoteesgh/go-library/internal/types"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	app *types.App
}

func NewAuthHandler(app *types.App) *AuthHandler {
	gob.Register(types.User{})
	return &AuthHandler{
		app: app,
	}
}

func (h *AuthHandler) ShowRegisterPage(w http.ResponseWriter, r *http.Request) {
	if tmpl, err := template.ParseFiles(
		"web/templates/register.html",
		h.app.Templates["guest"],
	); err != nil {
		panic(err)
	} else {
		tmpl.Execute(w, nil)
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	user := &types.User{
		Name:  r.FormValue("name"),
		Email: r.FormValue("email"),
	}

	password := hashPassword(r.FormValue("password"))

	_, err := h.app.DB.Exec(`
		INSERT INTO users (name, email, password)
		VALUES (?, ?, ?)
	`, user.Name, user.Email, password)

	if err != nil {
		fmt.Fprint(w, `<p class="text-red-700 font-semibold">Fail</p>`)
	} else {
		fmt.Fprint(w, `<p class="text-green-700 font-semibold">Success</p>`)
	}
}

func (h *AuthHandler) ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	if tmpl, err := template.ParseFiles(
		"web/templates/login.html",
		h.app.Templates["guest"],
	); err != nil {
		panic(err)
	} else {
		tmpl.Execute(w, nil)
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var password string

	user := &types.User{
		Email: r.FormValue("email"),
	}

	rows, err := h.app.DB.Query(`
		SELECT
			id,
			name,
			email,
			password,
			created_at,
			updated_at
		FROM users
		WHERE email = ?
	`, user.Email)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		if err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&password,
			&user.Created_at,
			&user.Updated_at,
		); err != nil {
			panic(err)
		}
	} else {
		fmt.Fprintf(w, `<p class="text-red-700 font-semibold">Your details are incorrect. Please try again.</p>`)
		return
	}

	if checkPasswordHash(r.FormValue("password"), password) {
		session, err := h.app.Session.Get(r, "auth")
		if err != nil {
			panic(err)
		}

		session.Values["user"] = user
		if err := session.Save(r, w); err != nil {
			panic(err)
		}

		w.Header().Set("HX-Location", "/")
		fmt.Fprintf(w, `<p class="text-green-700 font-semibold">Success %#v</p>`, user)
	} else {
		fmt.Fprint(w, `<p class="text-red-700 font-semibold">Your details are incorrect. Please try again.</p>`)
	}
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := h.app.Session.Get(r, "auth")
	if err != nil {
		panic(err)
	}

	session.Values["user"] = nil
	if err := session.Save(r, w); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
