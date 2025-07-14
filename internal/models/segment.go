package models

import (
	"database/sql"
	"time"

	"segment-service/internal/db"
)

// Segment описывает структуру сегмента пользователей.
type Segment struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	DistributionRatio float64   `json:"distribution_ratio"`
	CreatedAt         time.Time `json:"created_at"`
}

// CreateSegment создает новый сегмент в базе данных.
func CreateSegment(s Segment) error {
	_, err := db.DB.Exec(`
		INSERT INTO segments (name, description, distribution_ratio)
		VALUES (?, ?, ?)
	`, s.Name, s.Description, s.DistributionRatio)

	return err
}

// GetSegmentByName возвращает сегмент по его имени.
func GetSegmentByName(name string) (*Segment, error) {
	row := db.DB.QueryRow(`SELECT id, name, description, distribution_ratio, created_at FROM segments WHERE name = ?`, name)

	var s Segment
	err := row.Scan(&s.ID, &s.Name, &s.Description, &s.DistributionRatio, &s.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &s, err
}

// DeleteSegment удаляет сегмент по имени.
func DeleteSegment(name string) error {
	_, err := db.DB.Exec(`DELETE FROM segments WHERE name = ?`, name)
	return err
}

// UpdateSegment обновляет описание и коэффициент распределения сегмента по имени.
func UpdateSegment(name, description string, distributionRatio float64) error {
	_, err := db.DB.Exec(`
		UPDATE segments SET description = ?, distribution_ratio = ? WHERE name = ?
	`, description, distributionRatio, name)
	return err
}

// UpdateSegmentDistributionRatio обновляет только коэффициент распределения сегмента по имени.
func UpdateSegmentDistributionRatio(name string, distributionRatio float64) error {
	_, err := db.DB.Exec(`
		UPDATE segments SET distribution_ratio = ? WHERE name = ?
	`, distributionRatio, name)
	return err
}
