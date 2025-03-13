package main

import (
	// "database/sql"

	"fmt"
	"forum/config"
	"forum/database"
	"forum/handlers"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.InitDB("./database/forum.db")
	config.InitTemplate()
	config.InitRegex()

	forumux := http.NewServeMux()
	forumux.HandleFunc("/login", handlers.SwitchLogin)
	forumux.HandleFunc("/register", handlers.SwitchRegister)
	forumux.HandleFunc("/logout",handlers.LogoutHandler)

	forumux.HandleFunc("/like", handlers.AuthMidleware(handlers.LikeHandler))
	forumux.HandleFunc("/post", handlers.AuthMidleware(handlers.PostHandler))
	forumux.HandleFunc("/comment", handlers.AuthMidleware(handlers.CommentHandler))

	forumux.HandleFunc("/", handlers.RootHandler)
	forumux.HandleFunc("/static/", handlers.StaticHnadler)

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", forumux)

	// fmt.Println("Available templates:", temp.DefinedTemplates())
	fmt.Println(err)
}
