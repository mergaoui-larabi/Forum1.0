package database


import (
	"database/sql"
	"log"
)

func AddPost(db *sql.DB, userID int, content string) error {
	query := `INSERT INTO posts (user_id, content) VALUES (?, ?)`
	_, err := db.Exec(query, userID, content)
	if err != nil {
		log.Println("Error inserting post:", err)
		return err
	}
	return nil
}