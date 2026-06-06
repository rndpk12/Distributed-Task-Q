package models

import "time"

type Task struct {
	ID        string `gorm:"primaryKey"`
	Type      string
	Payload   string
	Status    string
	Retries   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
