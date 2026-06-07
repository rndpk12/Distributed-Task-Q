package ws

import (
	"encoding/json"
	"log"

	"github.com/rndpk/distributed-task-queue/internal/queue"
)

func StartSubscriber() {

	pubsub := queue.RDB.Subscribe(
		queue.Ctx,
		"events",
	)

	ch := pubsub.Channel()

	go func() {

		for msg := range ch {

			var event Event

			err := json.Unmarshal(
				[]byte(msg.Payload),
				&event,
			)

			if err != nil {
				log.Println(err)
				continue
			}

			log.Printf(
				"Received Redis event: %s - %s",
				event.Type,
				event.Message,
			)

			Broadcast(event)
		}
	}()
}
