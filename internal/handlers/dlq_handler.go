package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rndpk/distributed-task-queue/internal/queue"
)

func GetDLQ(c *gin.Context) {

	tasks, err := queue.RDB.LRange(
		queue.Ctx,
		"tasks:dlq",
		0,
		-1,
	).Result()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"dlq_tasks": tasks,
			"count":     len(tasks),
		},
	)
}
