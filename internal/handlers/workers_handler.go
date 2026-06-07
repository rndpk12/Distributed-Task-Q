package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rndpk/distributed-task-queue/internal/db"
	"github.com/rndpk/distributed-task-queue/internal/models"
)

func GetWorkers(c *gin.Context) {

	RefreshWorkerStatus()

	var workers []models.Worker

	if err := db.DB.
		Order("tasks_processed desc").
		Find(&workers).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, workers)
}
