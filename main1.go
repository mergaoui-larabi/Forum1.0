package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func PrintPosts(db *sql.DB) {
	rows, err := db.Query("SELECT id, user_id, content, created_at FROM post")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Posts in the database:")
	for rows.Next() {
		var id, userID int
		var content, createdAt string

		err := rows.Scan(&id, &userID, &content, &createdAt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, User ID: %d, Content: %s, Created At: %s\n", id, userID, content, createdAt)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func mainn() {
	db, err := sql.Open("sqlite3", "./database/forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	PrintPosts(db)
}
