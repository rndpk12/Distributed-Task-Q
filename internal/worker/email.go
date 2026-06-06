package worker

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/rndpk/distributed-task-queue/internal/models"
)

func ProcessEmail(task models.Task) error {

	log.Printf("Sending email: %s", task.Payload)

	if strings.Contains(task.Payload, "fail") {
		return fmt.Errorf("simulated email failure")
	}

	time.Sleep(2 * time.Second)

	log.Printf("Email sent: %s", task.Payload)

	return nil
}
