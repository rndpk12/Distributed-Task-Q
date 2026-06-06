package worker

import (
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/rndpk/distributed-task-queue/internal/db"
	"github.com/rndpk/distributed-task-queue/internal/metrics"
	"github.com/rndpk/distributed-task-queue/internal/models"
	"github.com/rndpk/distributed-task-queue/internal/queue"
	"github.com/rndpk/distributed-task-queue/internal/ws"
)

func Start() {

	workerID := uuid.New().String()[:8]

	log.Printf("Worker %s started\n", workerID)

	for {

		taskID, err := queue.RDB.BRPop(
			queue.Ctx,
			0,
			"tasks",
		).Result()

		if err != nil {
			log.Println(err)
			continue
		}

		id := taskID[1]

		var task models.Task

		if err := db.DB.First(&task, "id = ?", id).Error; err != nil {
			log.Println(err)
			continue
		}

		task.Status = "processing"
		db.DB.Save(&task)

		log.Printf("[%s] Processing task %s\n", workerID, task.ID)
		ws.Broadcast(ws.Event{
			Type:    "processing",
			Message: task.ID,
		})

		var processErr error

		switch task.Type {

		case "email":
			processErr = ProcessEmail(task)

		default:
			log.Printf("unknown task type %s", task.Type)
		}

		if processErr != nil {

			task.Retries++
			metrics.IncrementRetried()

			queue.RDB.Incr(
				queue.Ctx,
				"metrics:retried",
			)

			if task.Retries <= 3 {

				log.Printf(
					"[%s] Retry %d for task %s\n",
					workerID,
					task.Retries,
					task.ID,
				)

				task.Status = "retrying"

				db.DB.Save(&task)

				backoff := time.Duration(
					1<<(task.Retries-1),
				) * time.Second

				log.Printf(
					"[%s] Waiting %v before retry\n",
					workerID,
					backoff,
				)

				time.Sleep(backoff)

				queue.RDB.LPush(
					queue.Ctx,
					"tasks",
					task.ID,
				)

				continue
			}

			log.Printf(
				"[%s] Task permanently failed %s\n",
				workerID,
				task.ID,
			)
			ws.Broadcast(ws.Event{
				Type:    "failed",
				Message: task.ID,
			})

			queue.RDB.LPush(
				queue.Ctx,
				"tasks:dlq",
				task.ID,
			)

			metrics.IncrementFailed()

			queue.RDB.LPush(
				queue.Ctx,
				"tasks:dlq",
				task.ID,
			)

			metrics.IncrementFailed()

			queue.RDB.Incr(
				queue.Ctx,
				"metrics:failed",
			)
			task.Status = "failed"

			db.DB.Save(&task)

			continue
		}

		metrics.IncrementProcessed()

		metrics.IncrementProcessed()

		queue.RDB.Incr(
			queue.Ctx,
			"metrics:processed",
		)
		task.Status = "completed"
		db.DB.Save(&task)

		log.Printf("[%s] Completed task %s\n", workerID, task.ID)
		ws.Broadcast(ws.Event{
			Type:    "completed",
			Message: task.ID,
		})
	}
}
