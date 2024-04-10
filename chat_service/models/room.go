package models

import (
	pq "github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type Room struct {
	gorm.Model
	Joinees pq.Int64Array `gorm:"type:integer[]"`
	EndedAt time.Time
	RoomID  string
}

func (*Room) TableName() string {
	return "room"
}
