package model

import "time"

type HistoryLog struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	LogID     int       `json:"log_id"`
	LogType   string    `json:"log_type"`
	Log       string    `json:"log"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
