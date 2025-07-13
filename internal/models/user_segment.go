package models

import (
	"math/rand"
	"segment-service/internal/db"
	"time"
)

func AddUserToSegment(userID, segmentID int) error {
	_, err := db.DB.Exec(`
		INSERT OR IGNORE INTO user_segments (user_id, segment_id)
		VALUES (?, ?)
	`, userID, segmentID)
	return err
}

func RemoveUserFromSegment(userID, segmentID int) error {
	_, err := db.DB.Exec(`
		DELETE FROM user_segments
		WHERE user_id = ? AND segment_id = ?
	`, userID, segmentID)
	return err
}

func GetUserSegments(userID int) ([]string, error) {
	rows, err := db.DB.Query(`
		SELECT s.name
		FROM segments s
		JOIN user_segments us ON s.id = us.segment_id
		WHERE us.user_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var segments []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		segments = append(segments, name)
	}

	return segments, nil
}

func AssignSegmentRandomly(segmentName string, ratio float64) (int, error) {
	// Получаем сегмент по имени
	segment, err := GetSegmentByName(segmentName)
	if err != nil || segment == nil {
		return 0, err
	}

	// Обновляем distribution_ratio в базе данных
	err = UpdateSegmentDistributionRatio(segmentName, ratio)
	if err != nil {
		return 0, err
	}

	rows, err := db.DB.Query(`SELECT id FROM users`)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var userIDs []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
		userIDs = append(userIDs, id)
	}

	total := len(userIDs)
	if total == 0 {
		return 0, nil
	}

	n := int(float64(total) * ratio)
	if n == 0 {
		return 0, nil
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(userIDs), func(i, j int) {
		userIDs[i], userIDs[j] = userIDs[j], userIDs[i]
	})

	selected := userIDs[:n]
	count := 0
	for _, uid := range selected {
		err := AddUserToSegment(uid, segment.ID)
		if err == nil {
			count++
		}
	}
	return count, nil
}

func GetSegmentUsers(segmentName string) ([]int, error) {
	rows, err := db.DB.Query(`
		SELECT u.id
		FROM users u
		JOIN user_segments us ON u.id = us.user_id
		JOIN segments s ON us.segment_id = s.id
		WHERE s.name = ?
		ORDER BY u.id
	`, segmentName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userIDs []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		userIDs = append(userIDs, id)
	}

	return userIDs, nil
}
