package models

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	SenderID    string
	ReceiverID  string
	Content     string
	ContentType string    `default:"text/plain; charset=utf-8"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	RoomID      uint
}

func (*Message) TableName() string {
	return "message"
}
