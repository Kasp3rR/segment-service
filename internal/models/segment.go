package models

import (
	"database/sql"
	"time"

	"segment-service/internal/db"
)

type Segment struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	DistributionRatio float64   `json:"distribution_ratio"`
	CreatedAt         time.Time `json:"created_at"`
}

func CreateSegment(s Segment) error {
	_, err := db.DB.Exec(`
		INSERT INTO segments (name, description, distribution_ratio)
		VALUES (?, ?, ?)
	`, s.Name, s.Description, s.DistributionRatio)

	return err
}

func GetSegmentByName(name string) (*Segment, error) {
	row := db.DB.QueryRow(`SELECT id, name, description, distribution_ratio, created_at FROM segments WHERE name = ?`, name)

	var s Segment
	err := row.Scan(&s.ID, &s.Name, &s.Description, &s.DistributionRatio, &s.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &s, err
}

func DeleteSegment(name string) error {
	_, err := db.DB.Exec(`DELETE FROM segments WHERE name = ?`, name)
	return err
}

func UpdateSegment(name, description string, distributionRatio float64) error {
	_, err := db.DB.Exec(`
		UPDATE segments SET description = ?, distribution_ratio = ? WHERE name = ?
	`, description, distributionRatio, name)
	return err
}
