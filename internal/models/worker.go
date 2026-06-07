package models

import "time"

type Worker struct {
	ID             string    `gorm:"primaryKey" json:"id"`
	Status         string    `json:"status"`
	TasksProcessed int       `json:"tasks_processed"`
	LastSeen       time.Time `json:"last_seen"`
}
