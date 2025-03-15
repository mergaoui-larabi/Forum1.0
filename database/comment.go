package database

import "forum/models"

func GetComments() []models.Comment {
	rows, err := DB.Query("SELECT comment FROM comments ORDER BY created_at ASC")
	if err != nil {
		return nil
	}
	defer rows.Close()
	var comments []models.Comment
	for rows.Next() {
		var c models.Comment
		rows.Scan(&c.Content)
		comments = append(comments, c)
	}
	return comments
}

func AddComment(userID, postID int, content string) error {
	_, err := DB.Exec("INSERT INTO comments (user_id, post_id, comment) VALUES (?, ?, ?)", userID, postID, content)
	return err
}
