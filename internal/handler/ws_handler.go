package handler

import (
	"net/http"

	"github.com/TonChan8028/go-realtime-chat/internal/hub"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var roomManager = hub.NewRoomManager()

func WebSocketEcho(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		username = "anonymous"
	}

	room := r.URL.Query().Get("room")
	if room == "" {
		room = "general"
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	h := roomManager.GetRoom(room)
	client := hub.NewClient(h, conn, username)
	h.Register(client)

	go client.WritePump()
	go client.ReadPump()
}
