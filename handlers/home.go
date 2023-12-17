package handlers

import (
	"html/template"
	"net/http"
)

var guestTemplate = "views/templates/guest.html"
var authTemplate = "views/templates/auth.html"

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/home.html", guestTemplate)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}
