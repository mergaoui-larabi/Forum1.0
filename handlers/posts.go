package handlers

import (
	"log"
	"net/http"

	"forum/database"

	_ "github.com/mattn/go-sqlite3"
)


func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	UserId := 1
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	content := r.FormValue("content")
	if content == "" {
		return
	}

	_, err := database.Db.Exec("INSERT INTO post (user_id, content) VALUES (?, ?)", UserId, content)
	if err != nil {
		log.Print(err)
		http.Error(w, "Failed to insert post ", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/page", http.StatusSeeOther)
}