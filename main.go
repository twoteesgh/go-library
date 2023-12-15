package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", "elliott@(127.0.0.1:3306)/gowebexamples")
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/books/{author}/{title}", showBook).Methods("GET")

	http.ListenAndServe(":8008", r)
}

func showBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "You are reading %s, by %s", vars["title"], vars["author"])
}
