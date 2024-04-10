package models

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	SenderID    uint
	ReceiverID  uint
	Content     string
	ContentType string    `default:"text/plain; charset=utf-8"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func (*Message) TableName() string {
	return "message"
}
