package model

import (
	"database/sql"
	"time"
)

type Trackable struct {
	ID        int       `json:"id,omitempty"`
	Tool      string    `json:"tool"`               // e.g., VSCode, Chrome
	Metadata  string    `json:"metadata,omitempty"` // e.g., file path, URL
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Duration  int       `json:"duration,omitempty"` // Duration in seconds
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func LogTrackable(db *sql.DB, trackable Trackable) (int64, error) {
	duration := int(trackable.EndTime.Sub(trackable.StartTime).Seconds())
	trackable.Duration = duration
	query := `
		INSERT INTO Trackable (tool, metadata, start_time, end_time, duration)
		VALUES (?, ?, ?, ?, ?)`
	result, err := db.Exec(query, trackable.Tool, trackable.Metadata, trackable.StartTime, trackable.EndTime, trackable.Duration)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
