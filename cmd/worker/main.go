package main

import (
	"github.com/rndpk/distributed-task-queue/internal/db"
	"github.com/rndpk/distributed-task-queue/internal/queue"
	"github.com/rndpk/distributed-task-queue/internal/worker"
)

func main() {

	db.Connect()
	queue.Connect()

	worker.Start()
}
