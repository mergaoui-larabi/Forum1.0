package handlers

import (
	"log"
	"net/http"
	"fmt"
	"forum/database"
	_ "github.com/mattn/go-sqlite3"
)

func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("hnaaaa")
	user_id := r.Context().Value(userIDKey).(int)
	// user_id := 1
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	content := r.FormValue("content")
	fmt.Println(content)
	if content == "" {
		return
	}

	err := database.InsertPost(user_id, content)
	if err != nil {
		log.Print(err)
		http.Error(w, "Failed to insert post ", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}