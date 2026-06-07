package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rndpk/distributed-task-queue/internal/db"
	"github.com/rndpk/distributed-task-queue/internal/models"
	"github.com/rndpk/distributed-task-queue/internal/queue"
	"github.com/rndpk/distributed-task-queue/internal/ws"
)

func GetDLQ(c *gin.Context) {

	var tasks []models.Task

	if err := db.DB.
		Where("status = ?", "failed").
		Order("updated_at desc").
		Find(&tasks).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, tasks)
}

func RetryDLQTask(c *gin.Context) {

	id := c.Param("id")

	// Find the task — must be in failed status
	var task models.Task

	if err := db.DB.First(&task, "id = ? AND status = ?", id, "failed").Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "task not found in DLQ",
		})
		return
	}

	// Reset task back to pending state
	task.Status = "pending"
	task.Retries = 0

	if err := db.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update task: " + err.Error(),
		})
		return
	}

	// Remove task ID from the DLQ Redis list (LREM removes all occurrences)
	queue.RDB.LRem(queue.Ctx, "tasks:dlq", 0, task.ID)

	// Push back onto the main task queue
	queue.RDB.LPush(queue.Ctx, "tasks", task.ID)

	// Broadcast real-time event via WebSocket
	ws.Broadcast(ws.Event{
		Type:    "TASK_RETRIED",
		Message: task.ID,
	})

	// Also publish via Redis Pub/Sub
	queue.PublishEvent(queue.Event{
		Type:    "TASK_RETRIED",
		Message: task.ID,
	})

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"task_id": task.ID,
	})
}
