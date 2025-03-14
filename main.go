package main

import (
	"fmt"
	"log"
	"net/http"
	"forum/handlers"
	"forum/models"

	_ "github.com/mattn/go-sqlite3"
	database "forum/database"
)

func main() {
	db := database.Initdb()
	database.UserTable(db)
	database.PostTable(db)
	database.LikeAndDislikeTable(db)
	//database.DislikeTable(db)
	database.CommentTable(db)

	http.HandleFunc("/", handlers.ForumHandler)
	// http.HandleFunc("/", handlers.Login)
	// http.HandleFunc("/regist", handlers.Regist)
	// http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/like", models.ToggleLikeHandler)
	http.HandleFunc("/dislike", models.ToggleDislikeHandler)
	http.HandleFunc("/comment", models.CommentHandler)
	http.HandleFunc("/likes/count", models.LikesCountHandler)
	http.HandleFunc("/dislikes/count", models.DislikesCountHandler)
	http.HandleFunc("/comments", models.CommentsHandler)
	http.HandleFunc("/static/", handlers.StaticHnadler)
	http.HandleFunc("/add_post", handlers.AddPostHandler)
	http.HandleFunc("/add_comment", handlers.AddCommentHandler)
	http.HandleFunc("/add_like", handlers.AddLikesAndDislikes)

	// handlers.ShowPosts()

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
