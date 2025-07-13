package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./segment_service.db")
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("DB unreachable: %v", err)
	}

	if err := InitSchema(DB); err != nil {
		log.Fatalf("Schema init failed: %v", err)
	}

	log.Println("📦 Database connected and ready.")
}
