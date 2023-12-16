package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	setup()

	// Register application routes
	r := mux.NewRouter()
	r.HandleFunc("/database/books", createBooksTable).Methods("GET")
	r.HandleFunc("/books/{author}/{title}", showBook).Methods("GET")

	http.ListenAndServe(":8008", r)
}

func setup() {
	// Dotenv initialization
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Database initialization
	sqlConnectString := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	sqlDb, sqlErr := sql.Open("mysql", sqlConnectString)
	if sqlErr != nil {
		panic(sqlErr)
	}
	if err := sqlDb.Ping(); err != nil {
		panic(err)
	}
	db = sqlDb
}

func createBooksTable(w http.ResponseWriter, r *http.Request) {
	query := `
		CREATE TABLE books (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			author VARCHAR(255) NOT NULL,
			title VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, "Books table created")
	}

}

func showBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "You are reading %s, by %s", vars["title"], vars["author"])
}
