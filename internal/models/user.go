package models

import "segment-service/internal/db"

type User struct {
	ID int `json:"id"`
}

func CreateUser(userID int) error {
	_, err := db.DB.Exec(`INSERT OR IGNORE INTO users (id) VALUES (?)`, userID)
	return err
}

func UserExists(userID int) (bool, error) {
	row := db.DB.QueryRow(`SELECT id FROM users WHERE id = ?`, userID)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return false, nil
	}
	return true, nil
}
