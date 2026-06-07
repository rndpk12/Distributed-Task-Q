package models

import "time"

type Worker struct {
	ID             string `gorm:"primaryKey"`
	Status         string
	TasksProcessed int
	LastSeen       time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
