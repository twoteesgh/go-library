package repositories

import "github.com/twoteesgh/go-library/internal/types"

type BookRepository struct {
	app *types.App
}

func CreateBookRepository(app *types.App) *BookRepository {
	return &BookRepository{
		app: app,
	}
}

func (r *BookRepository) CreateBook(title string, author string) (types.Book, error) {
    book := types.Book{
        Title: title,
        Author: author,
    }

	if _, err := r.app.DB.Exec(`
        INSERT INTO books (title, author)
        VALUES (?, ?)
    `, book.Title, book.Author); err != nil {
        return types.Book{}, err
	}

    return book, nil
}
