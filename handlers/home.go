package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

var guestTemplate = "views/templates/guest.html"
var authTemplate = "views/templates/auth.html"

type Book struct {
	title  string
	author string
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/home.html", guestTemplate)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}

func CreateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("bookTitle")
		author := r.FormValue("bookAuthor")
		_, err := db.Exec(`
			INSERT INTO books (title, author)
			VALUES (?, ?)
		`, title, author)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "<li>%s - %s</li>", title, author)
	}
}

func GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
			SELECT title, author FROM books LIMIT 50	
		`)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		var books []Book
		for rows.Next() {
			var book Book
			if err := rows.Scan(
				&book.title,
				&book.author,
			); err != nil {
				panic(err)
			}
			books = append(books, book)
		}

		for _, book := range books {
			fmt.Fprintf(w, "<li>%s - %s</li>", book.title, book.author)
		}
	}
}
