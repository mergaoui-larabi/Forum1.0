package models

import (
	"encoding/json"
	"fmt"
	"forum/database"
	"net/http"
)

func addComment(userID, postID int, content string) error {
	_, err := database.Db.Exec("INSERT INTO comments (user_id, post_id, content) VALUES (?, ?, ?)", userID, postID, content)
	return err
}

type Comment struct {
	Content string `json:"content"`
}

func getComments() []Comment {
	rows, err := database.Db.Query("SELECT content FROM comments ORDER BY created_at ASC")
	if err != nil {
		return nil
	}
	defer rows.Close()
	var comments []Comment
	for rows.Next() {
		var c Comment
		rows.Scan(&c.Content)
		comments = append(comments, c)
	}
	return comments
}

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		PostID  int    `json:"post_id"`
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := addComment(1, req.PostID, req.Content); err != nil {
		http.Error(w, "Failed to add comment", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Comment added successfully!")
}

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getComments())
}
