package models

import "segment-service/internal/db"

// User описывает структуру пользователя.
type User struct {
	ID int `json:"id"`
}

// CreateUser создает пользователя с заданным ID.
func CreateUser(userID int) error {
	_, err := db.DB.Exec(`INSERT OR IGNORE INTO users (id) VALUES (?)`, userID)
	return err
}

// UserExists проверяет, существует ли пользователь с заданным ID.
func UserExists(userID int) (bool, error) {
	row := db.DB.QueryRow(`SELECT id FROM users WHERE id = ?`, userID)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return false, nil
	}
	return true, nil
}
