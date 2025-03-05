package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/database"
	"forum/handlers"
	"forum/models"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := database.Initdb()
	database.UserTable(db)
	database.PostTable(db)
	database.LikeAndDislikeTable(db)
	database.DislikeTable(db)
	database.CommentTable(db)

	http.HandleFunc("/page", handlers.ForumHandler)
	http.HandleFunc("/", handlers.Login)
	http.HandleFunc("/regist", handlers.Regist)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/like", models.ToggleLikeHandler)
	http.HandleFunc("/dislike", models.ToggleDislikeHandler)
	http.HandleFunc("/comment", models.CommentHandler)
	http.HandleFunc("/likes/count", models.LikesCountHandler)
	http.HandleFunc("/dislikes/count", models.DislikesCountHandler)
	http.HandleFunc("/comments", models.CommentsHandler)
	http.HandleFunc("/static/", handlers.StaticHnadler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
