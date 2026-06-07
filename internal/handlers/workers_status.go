package handlers

import (
	"time"

	"github.com/rndpk/distributed-task-queue/internal/db"
	"github.com/rndpk/distributed-task-queue/internal/models"
)

func RefreshWorkerStatus() {

	var workers []models.Worker

	db.DB.Find(&workers)

	for _, worker := range workers {

		if time.Since(worker.LastSeen) > 30*time.Second {

			db.DB.Model(&worker).
				Update("status", "offline")
		}
	}
}
