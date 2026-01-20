package hub

import (
	"log"
	"time"

	"github.com/TonChan8028/go-realtime-chat/internal/model"
	"github.com/gorilla/websocket"
)

type Client struct {
	hub      *Hub
	conn     *websocket.Conn
	send     chan model.Message
	username string
}

func NewClient(hub *Hub, conn *websocket.Conn, username string) *Client {
	return &Client{
		hub:      hub,
		conn:     conn,
		send:     make(chan model.Message),
		username: username,
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		var msg model.Message
		if err := c.conn.ReadJSON(&msg); err != nil {
			log.Println("read error:", err)
			break
		}
		msg.Type = model.MessageTypeMessage
		msg.Username = c.username
		msg.Timestamp = time.Now()

		c.hub.Broadcast(msg)
	}
}

func (c *Client) WritePump() {
	defer c.conn.Close()

	for msg := range c.send {
		if err := c.conn.WriteJSON(msg); err != nil {
			log.Println("write error:", err)
			break
		}
	}
}
