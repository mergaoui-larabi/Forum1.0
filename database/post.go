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

func InsertPost(user_id int, content string) error {
	_, err := DB.Exec("INSERT INTO post (user_id, content) VALUES (?, ?)", user_id, content)
	return err
}
