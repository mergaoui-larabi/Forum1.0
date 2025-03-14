package handlers

import (
	"forum/database"
	"log"
	"net/http"
	"fmt"
)

func AddLikesAndDislikes(w http.ResponseWriter, r *http.Request){
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
	var likeOrDislike bool
	if r.FormValue("value") == "1"{
		likeOrDislike = true
	}else{
		likeOrDislike = false
	}
	//check ila l user deja reacta
	var check int
	err := database.Db.QueryRow("SELECT COUNT(1) FROM likes_dislikes WHERE user_id = ? AND post_id = ?", UserId, post_id).Scan(&check)
	if err != nil {
		log.Println("Error checking like existence:", err)
	}
	fmt.Println("Cheeeeeck iiiss :",check)
	if check > 0 {
		_, err = database.Db.Exec("UPDATE likes_dislikes SET is_like = ? WHERE user_id = ? AND post_id = ?", likeOrDislike, UserId, post_id)
		fmt.Println("Value is ",likeOrDislike)
	} else {
		if likeOrDislike{
			_, err := database.Db.Exec("INSERT INTO likes_dislikes (user_id, post_id, is_like) VALUES (?, ?, ?)", UserId,post_id, likeOrDislike)
			if err != nil {
				log.Print(err)
				http.Error(w, "Failed to insert like", http.StatusInternalServerError)
				return
			}
		}else {
			_, err := database.Db.Exec("INSERT INTO likes_dislikes (user_id, post_id, is_like) VALUES (?, ?, ?)", UserId,post_id, likeOrDislike)
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