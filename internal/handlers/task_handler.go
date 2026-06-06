package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/rndpk/distributed-task-queue/internal/db"
	"github.com/rndpk/distributed-task-queue/internal/models"
	"github.com/rndpk/distributed-task-queue/internal/queue"
)

type CreateTaskRequest struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

func CreateTask(c *gin.Context) {

	var req CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	task := models.Task{
		ID:      uuid.New().String(),
		Type:    req.Type,
		Payload: req.Payload,
		Status:  "pending",
	}

	db.DB.Create(&task)

	queue.RDB.LPush(
		queue.Ctx,
		"tasks",
		task.ID,
	)

	c.JSON(
		http.StatusCreated,
		task,
	)
}
