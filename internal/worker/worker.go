package worker

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/rndpk/distributed-task-queue/internal/db"
	"github.com/rndpk/distributed-task-queue/internal/metrics"
	"github.com/rndpk/distributed-task-queue/internal/models"
	"github.com/rndpk/distributed-task-queue/internal/queue"
)

func Start() {

	workerID := uuid.New().String()[:8]

	log.Printf("Worker %s started\n", workerID)

	db.DB.Create(&models.Worker{
		ID:             workerID,
		Status:         "active",
		TasksProcessed: 0,
		LastSeen:       time.Now(),
	})

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

		db.DB.Model(&models.Worker{}).
			Where("id = ?", workerID).
			Updates(map[string]interface{}{
				"last_seen": time.Now(),
			})

		log.Printf("[%s] Processing task %s\n", workerID, task.ID)
		queue.PublishEvent(queue.Event{
			Type:    "TASK_PROCESSING",
			Message: task.ID,
		})

		var processErr error

		switch task.Type {

		case "email":
			processErr = ProcessEmail(task)

		case "fail":
			log.Println("FAIL TASK HIT")
			processErr = errors.New("intentional failure")

		default:
			log.Printf(
				"unknown task type %s",
				task.Type,
			)
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
			queue.PublishEvent(queue.Event{
				Type:    "TASK_FAILED",
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

		db.DB.Model(&models.Worker{}).
			Where("id = ?", workerID).
			Update(
				"tasks_processed",
				gorm.Expr("tasks_processed + ?", 1),
			)

		log.Printf("[%s] Completed task %s\n", workerID, task.ID)
		queue.PublishEvent(queue.Event{
			Type:    "TASK_COMPLETED",
			Message: task.ID,
		})
	}
}
