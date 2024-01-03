package handlers

import (
	"fmt"
	"net/http"

	"github.com/twoteesgh/go-library/internal/types"
)

type BookHandler struct {
	app *types.App
}

type Book struct {
	title  string
	author string
}

func NewBookHandler(app *types.App) *BookHandler {
	return &BookHandler{
		app: app,
	}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("bookTitle")
	author := r.FormValue("bookAuthor")

	// TODO: Make a repository
	if _, err := h.app.DB.Exec(`
		INSERT INTO books (title, author)
		VALUES (?, ?)
	`, title, author); err != nil {
		panic(err)
	}

	fmt.Fprint(w, h.app.GetHtml("components/books/list-item", map[string]string{
		"title":  title,
		"author": author,
	}))
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	// TODO: Make a repository
	rows, err := h.app.DB.Query(`
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
		fmt.Fprint(w, h.app.GetHtml("components/books/list-item", map[string]string{
			"title":  book.title,
			"author": book.author,
		}))
	}
}
