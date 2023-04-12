package models

import (
	"time"
)

type Goal struct {
	ID          uint64    `json:"id"`
	UserID      uint64    `json:"user_id"`
	Name        string    `json:"name"`
	ProjectID   uint64    `json:"project_id"`
	HoursCount  float64   `json:"hours_count"`
	Description string    `json:"description"`
	TimeStart   time.Time `json:"time_start"`
	TimeEnd     time.Time `json:"time_end"`
}
