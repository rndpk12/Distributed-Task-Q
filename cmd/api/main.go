package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rndpk/distributed-task-queue/internal/db"
	"github.com/rndpk/distributed-task-queue/internal/handlers"
	"github.com/rndpk/distributed-task-queue/internal/queue"
)

func main() {

	db.Connect()
	queue.Connect()

	r := gin.Default()

	r.POST(
		"/tasks",
		handlers.CreateTask,
	)

	r.GET(
		"/metrics",
		handlers.GetMetrics,
	)

	r.GET(
		"/dlq",
		handlers.GetDLQ,
	)

	r.GET(
		"/ws",
		handlers.WebSocketHandler,
	)

	r.StaticFile(
		"/dashboard",
		"./web/dashboard.html",
	)

	r.Run(":8081")
}
