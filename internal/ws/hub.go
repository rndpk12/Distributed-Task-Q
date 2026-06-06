package ws

import (
	"github.com/gorilla/websocket"
)

var Clients = make(map[*websocket.Conn]bool)

type Event struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func Broadcast(event Event) {

	for client := range Clients {

		err := client.WriteJSON(event)

		if err != nil {

			client.Close()

			delete(
				Clients,
				client,
			)
		}
	}
}
