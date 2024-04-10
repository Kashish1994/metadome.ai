package helper

import "github.com/gorilla/websocket"

type Message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Content  string `json:"content"`
	Type     string `json:"type"` // Text/Image/Audio/Video --> For now supporting just text messages
}

type Room struct {
	SenderUsername   string `json:"sender_username"`
	ReceiverUsername string `json:"receiver_username"`
	RoomID           string `json:"room_id"`
	Broadcaster      chan *Message
	Connections      map[string]*websocket.Conn
}
