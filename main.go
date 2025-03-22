package main

import (
	// "database/sql"

	"fmt"
	"forum/config"
	"forum/database"
	"forum/handlers"
	"net/http"

	// "forum/models"

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
	forumux.HandleFunc("/login", handlers.SwitchLogin)
	forumux.HandleFunc("/register", handlers.SwitchRegister)
	forumux.HandleFunc("/logout", handlers.LogoutHandler)

	forumux.HandleFunc("/profile", handlers.AuthMidleware(handlers.ProfilHandler))
	forumux.HandleFunc("/profile/update/{value}", handlers.AuthMidleware(handlers.UpddateProfile))
	forumux.HandleFunc("/profile/update/{value}/save", handlers.AuthMidleware(handlers.SaveChanges))
	forumux.HandleFunc("/profile/delete", handlers.AuthMidleware(handlers.ServeDelete))
	forumux.HandleFunc("/profile/delete/confirm", handlers.AuthMidleware(handlers.DeleteConfirmation))
	// fmt.Println("server is running")
	forumux.HandleFunc("/", handlers.RootHandler)
	forumux.HandleFunc("/static/", handlers.StaticHnadler)
	forumux.HandleFunc("/like", database.ToggleLikeHandler)
	forumux.HandleFunc("/dislike", database.ToggleDislikeHandler)
	forumux.HandleFunc("/comment", database.CommentHandler)
	forumux.HandleFunc("/likes/count", database.LikesCountHandler)
	forumux.HandleFunc("/dislikes/count", database.DislikesCountHandler)
	forumux.HandleFunc("/comments", database.CommentsHandler)
	forumux.HandleFunc("/add_post", handlers.AddPostHandler)
	forumux.HandleFunc("/add_comment", handlers.AddCommentHandler)
	forumux.HandleFunc("/add_like", handlers.AddLikesAndDislikes)

	fmt.Println("Server running on ", SERVERURL)
	err := http.ListenAndServe(PORT, forumux)

	// fmt.Println("Available templates:", temp.DefinedTemplates())
	fmt.Println(err)
}
