package handlers

import (
	"fmt"
	"forum/database"
	"log"
	"net/http"
)

func AddLikesAndDislikes(w http.ResponseWriter, r *http.Request) {
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
	var likeOrDislike bool
	if r.FormValue("value") == "1" {
		likeOrDislike = true
	} else {
		likeOrDislike = false
	}
	//check ila l user deja reacta
	check := database.IsReacted(user_id,post_id)
	if check {
		_, err = database.Db.Exec("UPDATE likes_dislikes SET is_like = ? WHERE user_id = ? AND post_id = ?", likeOrDislike, user_id, post_id)
		fmt.Println("Value is ", likeOrDislike)
	} else {
		if likeOrDislike {
			_, err := database.Db.Exec("INSERT INTO likes_dislikes (user_id, post_id, is_like) VALUES (?, ?, ?)", user_id, post_id, likeOrDislike)
			if err != nil {
				log.Print(err)
				http.Error(w, "Failed to insert like", http.StatusInternalServerError)
				return
			}
		} else {
			_, err := database.Db.Exec("INSERT INTO likes_dislikes (user_id, post_id, is_like) VALUES (?, ?, ?)", user_id, post_id, likeOrDislike)
			if err != nil {
				log.Print(err)
				http.Error(w, "Failed to insert dislikeeeee", http.StatusInternalServerError)
				return
			}
		}
	}
	// fmt.Println("Inserting like/dislike:", post_id, likeOrDislike)
	// fmt.Println(r.FormValue("value"))
	http.Redirect(w, r, "/page", http.StatusSeeOther)

}
