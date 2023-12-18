package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/twoteesgh/go-library/handlers"
	"github.com/twoteesgh/go-library/services"
)

type Library struct {
	home  func(http.ResponseWriter, *http.Request)
	books *handlers.BookHandler
	auth  *handlers.AuthHandler
}

func main() {
	library := setup()

	// Register application routes
	r := mux.NewRouter()
	r.HandleFunc("/", library.home).Methods("GET")

	// Auth routes
	r.HandleFunc("/register", library.auth.ShowRegisterPage).Methods("GET")
	r.HandleFunc("/register", library.auth.Register).Methods("POST")
	r.HandleFunc("/login", library.auth.ShowLoginPage).Methods("GET")
	r.HandleFunc("/login", library.auth.Login).Methods("POST")

	// Book routes
	r.HandleFunc("/books", library.books.CreateBook).Methods("POST")
	r.HandleFunc("/books", library.books.GetBooks).Methods("GET")

	// FS initialization
	fs := http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", fs)).
		Methods("GET")

	http.ListenAndServe(":8008", r)
}

func setup() *Library {
	// Dotenv initialization
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db := services.DB()
	app := services.NewAppService(db)

	// Return app handlers
	return &Library{
		home:  app.Home,
		books: handlers.NewBookHandler(app),
		auth:  handlers.NewAuthHandler(app),
	}
}
