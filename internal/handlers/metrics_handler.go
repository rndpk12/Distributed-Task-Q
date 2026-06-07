package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rndpk/distributed-task-queue/internal/queue"
)

func GetMetrics(c *gin.Context) {

	processed, _ := queue.RDB.Get(
		queue.Ctx,
		"metrics:processed",
	).Result()

	failed, _ := queue.RDB.Get(
		queue.Ctx,
		"metrics:failed",
	).Result()

	retried, _ := queue.RDB.Get(
		queue.Ctx,
		"metrics:retried",
	).Result()

	queueDepth, _ := queue.RDB.LLen(
		queue.Ctx,
		"tasks",
	).Result()

	c.JSON(http.StatusOK, gin.H{
		"processed": processed,
		"failed":    failed,
		"retried":   retried,
		"queue":     queueDepth,
	})

}
