package db

import (
	"database/sql"
	"log"
)

func InitSchema(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY
	);

	CREATE TABLE IF NOT EXISTS segments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL,
		description TEXT,
		distribution_ratio REAL DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS user_segments (
		user_id INTEGER,
		segment_id INTEGER,
		PRIMARY KEY (user_id, segment_id),
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (segment_id) REFERENCES segments(id) ON DELETE CASCADE
	);
	`

	_, err := db.Exec(schema)
	if err != nil {
		log.Printf("Error creating schema: %v", err)
		return err
	}

	log.Println("Database schema initialized")
	return nil
}
