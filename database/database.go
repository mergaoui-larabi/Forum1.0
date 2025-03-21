package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var DB *sql.DB

func InitDB(datasource string) *sql.DB {
	var err error
	DB, err = sql.Open("sqlite3", datasource)
	if err != nil {
		log.Fatalf("failed to open the database: %v", err)
	}
	err = CreateTable(DB)
	if err != nil {
		log.Fatalf("failed to create tables: %v", err)
	}
	return DB
}

func CreateTable(db *sql.DB) error {
	schema, err := os.ReadFile("./database/schema.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(schema))
	return err
}

func CloseDB() {
	err := DB.Close()
	fmt.Println("closing the databse")
	if err != nil {
		log.Fatal("Error closing database:", err)
	}
}


