package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rndpk/distributed-task-queue/internal/db"
	"github.com/rndpk/distributed-task-queue/internal/models"
)

func GetTasks(c *gin.Context) {

	var tasks []models.Task

	if err := db.DB.
		Order("created_at desc").
		Find(&tasks).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(
		http.StatusOK,
		tasks,
	)
}
