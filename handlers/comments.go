package handlers

import (

	// "go/doc/comment"
	"forum/database"
	"log"

	// "fmt"
	"net/http"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(userIDKey).(int)
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	post_id := r.FormValue("post_id")
	if post_id == "" {
		http.Error(w, "Missing post ID", http.StatusBadRequest)
		return
	}

	comment := r.FormValue("content")
	if comment == "" {
		return
	}

	err := database.AddComment(user_id, post_id, comment)
	if err != nil {
		log.Print(err)
		http.Error(w, "Failed to insert comment", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/page", http.StatusSeeOther)
}
