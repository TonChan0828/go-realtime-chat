package hub

import (
	"go-realtime-chat/internal/model"
	"time"

	"github.com/TonChan8028/go-realtime-chat/internal/model"
)

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan model.Message
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan model.Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *Hub) Register(client *Client) {
	h.register <- client

	h.broadcast <- model.Message{
		Type:      model.MessageTypeJoin,
		Username:  client.username,
		Content:   "joined",
		Timestamp: time.Now(),
	}
}

func (h *Hub) Unregister(client *Client) {
	h.unregister <- client

	h.broadcast <- model.Message{
		Type:      model.MMessageTypeLeave,
		Username:  client.username,
		Content:   "left",
		Timestamp: time.Now(),
	}
}

func (h *Hub) Broadcast(msg model.Message) {
	h.broadcast <- msg
}
