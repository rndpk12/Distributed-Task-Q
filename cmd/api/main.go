package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rndpk/distributed-task-queue/internal/db"
	"github.com/rndpk/distributed-task-queue/internal/handlers"
	"github.com/rndpk/distributed-task-queue/internal/queue"
	"github.com/rndpk/distributed-task-queue/internal/ws"
)

func main() {

	db.Connect()
	queue.Connect()

	ws.StartSubscriber()

	r := gin.Default()

	r.POST(
		"/tasks",
		handlers.CreateTask,
	)

	r.GET(
		"/tasks",
		handlers.GetTasks,
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

	r.GET(
		"/workers",
		handlers.GetWorkers,
	)

	r.StaticFile(
		"/dashboard",
		"./web/dashboard.html",
	)

	r.Run(":8081")
}
