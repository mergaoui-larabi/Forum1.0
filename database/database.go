package database

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func Initdb() *sql.DB {
	var err error
	Db, err = sql.Open("sqlite3", "./database/forum.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	return Db
}

func UserTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		hashed_password TEXT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("SQL Error: %v\nQuery: %s", err, query)
		log.Fatal("Error creating users table")
	}
}

func PostTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		content TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}


func LikeAndDislikeTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS likes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Error creating likes table:", err)
	}
}

func CommentTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		comment TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Error creating comments table:", err)
	}
}

func CategoriesTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Error creating categories table:", err)
	}
}
