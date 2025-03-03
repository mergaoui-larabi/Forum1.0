package database

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func Initdb() {
	var err error
	Db, err = sql.Open("sqlite3", "./database/forum.db")
	if err != nil {
		log.Fatal(err)
	}

	createLikesTable := `
	CREATE TABLE IF NOT EXISTS likes (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    user_id INTEGER NOT NULL,
	    post_id INTEGER NOT NULL,
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = Db.Exec(createLikesTable)
	if err != nil {
		log.Fatal("Error creating likes table:", err)
	}

	createCommentsTable := `
	CREATE TABLE IF NOT EXISTS comments (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    user_id INTEGER NOT NULL,
	    post_id INTEGER NOT NULL,
	    content TEXT NOT NULL,
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = Db.Exec(createCommentsTable)
	if err != nil {
		log.Fatal("Error creating comments table:", err)
	}
}
