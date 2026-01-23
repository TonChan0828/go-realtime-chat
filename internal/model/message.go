package model

import "time"

type MessageType string

const (
	MessageTypeJoin    MessageType = "join"
	MessageTypeLeave   MessageType = "leave"
	MessageTypeMessage MessageType = "message"
	MessageTypeSystem  MessageType = "system"
)

type Message struct {
	Type      MessageType `json:"type"`
	Username  string      `json:"username"`
	Content   string      `json:"content"`
	Timestamp time.Time   `json:"timestamp"`
}
