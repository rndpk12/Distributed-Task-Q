package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/rndpk/distributed-task-queue/internal/ws"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(c *gin.Context) {

	conn, err := upgrader.Upgrade(
		c.Writer,
		c.Request,
		nil,
	)

	if err != nil {
		return
	}

	ws.Clients[conn] = true

	log.Printf(
		"WebSocket connected. Clients: %d",
		len(ws.Clients),
	)

	for {

		_, _, err := conn.ReadMessage()

		if err != nil {

			delete(ws.Clients, conn)

			conn.Close()

			log.Printf(
				"WebSocket disconnected. Clients: %d",
				len(ws.Clients),
			)

			break
		}
	}
}
