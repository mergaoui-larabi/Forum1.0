package main

import (
	// "database/sql"

	"fmt"
	"forum/config"
	"forum/database"
	"forum/handlers"
	auth "forum/handlers/authentification"
	post "forum/handlers/posts"
	static "forum/handlers/static"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const (
	PORT      = ":8080"
	SERVERURL = "http://localhost:8080"
)

func main() {
	database.InitDB("./database/forum.db")
	config.InitTemplate()
	config.InitRegex()

	forumux := http.NewServeMux()
	forumux.HandleFunc("/login", auth.SwitchLogin)
	forumux.HandleFunc("/register", auth.SwitchRegister)
	forumux.HandleFunc("/logout", auth.LogoutHandler)

	forumux.HandleFunc("/profile", auth.AuthMidleware(auth.ProfilHandler))
	forumux.HandleFunc("/profile/update/{value}", auth.AuthMidleware(auth.UpddateProfile))
	forumux.HandleFunc("/profile/update/{value}/save", auth.AuthMidleware(auth.SaveChanges))
	forumux.HandleFunc("/profile/delete", auth.AuthMidleware(auth.ServeDelete))
	forumux.HandleFunc("/profile/delete/confirm", auth.AuthMidleware(auth.DeleteConfirmation))

	forumux.HandleFunc("/like", auth.AuthMidleware(post.LikeHandler)) // TODO : generate ur own routes
	forumux.HandleFunc("/post", auth.AuthMidleware(post.PostHandler))
	forumux.HandleFunc("/comment", auth.AuthMidleware(post.CommentHandler))

	forumux.HandleFunc("/", handlers.RootHandler)
	forumux.HandleFunc("/static/", static.StaticHandler)

	fmt.Println("Server running on ", SERVERURL)
	err := http.ListenAndServe(PORT, forumux)

	// fmt.Println("Available templates:", temp.DefinedTemplates())
	fmt.Println(err)
}
