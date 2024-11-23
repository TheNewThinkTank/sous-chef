package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var DB *sql.DB

func InitDB(path string) {
	var err error
	DB, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create tables if they don't exist
	if err := createTables(); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
}

func createTables() error {
	schema := `
	CREATE TABLE IF NOT EXISTS recipes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		ingredients TEXT,
		instructions TEXT
	);
	CREATE TABLE IF NOT EXISTS pantry (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ingredient_name TEXT NOT NULL,
		quantity REAL,
		unit TEXT,
		expiry_date TEXT
	);
	`
	_, err := DB.Exec(schema)
	return err
}
