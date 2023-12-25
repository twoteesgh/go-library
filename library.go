package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/twoteesgh/go-library/handlers"
	"github.com/twoteesgh/go-library/services"
)

type App struct {
	home  *handlers.HomeHandler
	books *handlers.BookHandler
	auth  *handlers.AuthHandler
}

func main() {
	app := setup()

	// Register application routes
	r := mux.NewRouter()
	r.HandleFunc("/", app.home.Home).Methods("GET")

	// Auth routes
	r.HandleFunc("/register", app.auth.ShowRegisterPage).Methods("GET")
	r.HandleFunc("/register", app.auth.Register).Methods("POST")
	r.HandleFunc("/login", app.auth.ShowLoginPage).Methods("GET")
	r.HandleFunc("/login", app.auth.Login).Methods("POST")

	// Book routes
	r.HandleFunc("/books", app.books.CreateBook).Methods("POST")
	r.HandleFunc("/books", app.books.GetBooks).Methods("GET")

	// FS initialization
	fs := http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", fs)).
		Methods("GET")

	http.ListenAndServe(":8008", r)
}

func setup() *App {
	// Dotenv initialization
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db := services.DB()
	app := services.NewAppService(db)

	// Return app handlers
	return &App{
		home:  handlers.NewHomeHandler(app),
		books: handlers.NewBookHandler(app),
		auth:  handlers.NewAuthHandler(app),
	}
}
