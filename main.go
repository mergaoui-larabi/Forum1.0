package main

import (
	// "database/sql"
	"fmt"
	"forum/config"
	"forum/handlers"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// var db *sql.DB

func main() {
	// initDB()
	// defer db.Close()

	config.InitTemplate()

	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/static/", handlers.StaticHnadler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	
	//auth
	http.HandleFunc("/login-form", handlers.FromHandler)
	http.HandleFunc("/register-form", handlers.FromHandler)
	// http.HandleFunc("/register", nil)
	// http.HandleFunc("/logout", nil)
	// http.HandleFunc("/authorized", nil)

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	// fmt.Println("Available templates:", temp.DefinedTemplates())
	fmt.Println(err)
}
