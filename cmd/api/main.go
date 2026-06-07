package main

import (
	"os"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.Default())

	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetTasks)
	r.GET("/metrics", handlers.GetMetrics)
	r.GET("/dlq", handlers.GetDLQ)
	r.POST("/dlq/retry/:id", handlers.RetryDLQTask)
	r.GET("/ws", handlers.WebSocketHandler)
	r.GET("/workers", handlers.GetWorkers)
	r.StaticFile("/dashboard", "./web/dashboard.html")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/dashboard")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	r.Run(":" + port)
}
