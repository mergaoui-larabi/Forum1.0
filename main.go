package main

import (
	// "database/sql"
	"fmt"
	"forum/handlers"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// var db *sql.DB

func main() {
	// initDB()
	// defer db.Close()

	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/static/", handlers.StaticHnadler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
