package handlers

import (
	"fmt"
	"net/http"

	"github.com/twoteesgh/go-library/internal/repositories"
	"github.com/twoteesgh/go-library/internal/types"
)

type BookHandler struct {
	app *types.App
    repo *repositories.BookRepository
}

func NewBookHandler(app *types.App) *BookHandler {
	return &BookHandler{
		app: app,
        repo: repositories.CreateBookRepository(app),
	}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("bookTitle")
	author := r.FormValue("bookAuthor")

    book, err := h.repo.CreateBook(title, author)

    if err != nil {
        panic(err)
    }

	fmt.Fprint(w, h.app.GetHtml("components/books/list-item", map[string]string{
		"title":  book.Title,
		"author": book.Author,
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

    var books []types.Book
	for rows.Next() {
		var book types.Book
		if err := rows.Scan(
			&book.Title,
			&book.Author,
		); err != nil {
			panic(err)
		}
		books = append(books, book)
	}

	for _, book := range books {
		fmt.Fprint(w, h.app.GetHtml("components/books/list-item", map[string]string{
			"title":  book.Title,
			"author": book.Author,
		}))
	}
}
