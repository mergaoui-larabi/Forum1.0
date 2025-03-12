package handlers

import (
	"forum/database"
	// "go/doc/comment"
	"log"

	// "fmt"
	"net/http"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	UserId := 1
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	post_id := r.FormValue("post_id")
	if post_id == "" {
		http.Error(w, "Missing post ID", http.StatusBadRequest)
		return
	}

	comment:= r.FormValue("content")
	if comment == "" {
		return
	}
	

	_, err := database.Db.Exec("INSERT INTO comments (user_id, post_id, comment) VALUES (?, ?, ?)", UserId,post_id, comment)
	if err != nil {
		log.Print(err)
		http.Error(w, "Failed to insert cooment", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/page", http.StatusSeeOther)
}