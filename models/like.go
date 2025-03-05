package models

import (
	"encoding/json"
	"log"
	"net/http"

	"forum/database"
)

func checkIfLiked(userID, postID int) (bool, error) {
	var count int
	err := database.Db.QueryRow("SELECT COUNT(*) FROM likes WHERE user_id = ? AND post_id = ?", userID, postID).Scan(&count)
	if err != nil {
		return false, err
	}
	
	return count > 0, nil
}

func ToggleLikeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := 1
	postID := 1

	liked, err := checkIfLiked(userID, postID)
	if err != nil {
		http.Error(w, "Error checking like status", http.StatusInternalServerError)
		return
	}

	if liked {
		_, err = database.Db.Exec("DELETE FROM likes WHERE user_id = ? AND post_id = ?", userID, postID)
		if err != nil {
			http.Error(w, "Failed to unlike post", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Post unliked successfully"))
	} else {
		_, err = database.Db.Exec("INSERT INTO likes (user_id, post_id) VALUES (?, ?)", userID, postID)
		if err != nil {
			http.Error(w, "Failed to like post", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Post liked successfully"))
	}
}

func getLikesCount(postID int) int {
	var count int
	err := database.Db.QueryRow("SELECT COUNT(*) FROM likes WHERE post_id = ?", postID).Scan(&count)
	if err != nil {
		log.Printf("Error fetching likes count: %v", err)
		return 0
	}
	return count
}

func LikesCountHandler(w http.ResponseWriter, r *http.Request) {
	postID := 1
	count := getLikesCount(postID)
	json.NewEncoder(w).Encode(map[string]int{"count": count})
}
func checkIfDisliked(userID, postID int) (bool, error) {
	var count int
	err := database.Db.QueryRow("SELECT COUNT(*) FROM dislikes WHERE user_id = ? AND post_id = ?", userID, postID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func ToggleDislikeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := 1
	postID := 1 

	disliked, err := checkIfDisliked(userID, postID)
	if err != nil {
		http.Error(w, "Error checking dislike status", http.StatusInternalServerError)
		return
	}

	if disliked {
		_, err = database.Db.Exec("DELETE FROM dislikes WHERE user_id = ? AND post_id = ?", userID, postID)
		if err != nil {
			http.Error(w, "Failed to remove dislike", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Dislike removed successfully"))
	} else {
		_, err = database.Db.Exec("INSERT INTO dislikes (user_id, post_id) VALUES (?, ?)", userID, postID)
		if err != nil {
			http.Error(w, "Failed to add dislike", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Dislike added successfully"))
	}
}

func getDislikesCount(postID int) int {
	var count int
	err := database.Db.QueryRow("SELECT COUNT(*) FROM dislikes WHERE post_id = ?", postID).Scan(&count)
	if err != nil {
		log.Printf("Error fetching dislikes count: %v", err)
		return 0
	}
	return count
}

func DislikesCountHandler(w http.ResponseWriter, r *http.Request) {
	postID := 1
	count := getDislikesCount(postID)
	json.NewEncoder(w).Encode(map[string]int{"count": count})
}