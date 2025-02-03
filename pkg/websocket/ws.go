package websocket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

func WsPage(c *gin.Context) {
	res := c.Writer
	req := c.Request

	// Попытка установить WebSocket соединение
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}).Upgrade(res, req, nil)

	if err != nil {
		log.Printf("Error upgrading connection: %v", err) // Логируем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not upgrade connection"})
		return
	}

	client := &Client{id: uuid.NewV4().String(), socket: conn, send: make(chan []byte)}

	Manager.register <- client

	go client.read()
	go client.write()
}
