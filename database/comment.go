package database

import (
	"encoding/json"
	"net/http"
)

func addComment(userID, postID int, content string) error {
	_, err := DB.Exec("INSERT INTO comments (user_id, post_id, comment) VALUES (?, ?, ?)", userID, postID, content)
	return err
}

type Comment struct {
	Content string `json:"content"`
}

func getComments() []Comment {
	rows, err := DB.Query("SELECT comment FROM comments ORDER BY created_at ASC")
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
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	err = addComment(1, req.PostID, req.Content)
	if err != nil {
		http.Error(w, "Failed to add comment", http.StatusInternalServerError)
		return
	}
}

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getComments())
}

func AddComment(user_id int, post_id string, comment string) error {
	_, err := DB.Exec("INSERT INTO comments (user_id, post_id, comment) VALUES (?, ?, ?)", user_id, post_id, comment)
	return err
}

