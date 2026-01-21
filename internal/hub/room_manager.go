package hub

import "sync"

type RoomManager struct {
	mu    sync.Mutex
	rooms map[string]*Hub
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms: make(map[string]*Hub),
	}
}

func (rm *RoomManager) GetRoom(name string) *Hub {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if hub, ok := rm.rooms[name]; ok {
		return hub
	}

	hub := NewHub()
	rm.rooms[name] = hub
	go hub.Run()

	return hub
}
