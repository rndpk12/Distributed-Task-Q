package queue

import (
	"encoding/json"
)

type Event struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func PublishEvent(event Event) error {

	data, err := json.Marshal(event)

	if err != nil {
		return err
	}

	return RDB.Publish(
		Ctx,
		"events",
		data,
	).Err()
}
