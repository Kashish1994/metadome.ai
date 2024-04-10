package models

import (
	"gorm.io/gorm"
	"time"
)

type Room struct {
	gorm.Model
	Joinees []int `gorm:"type:integer[]"`
	EndedAt time.Time
}

func (*Room) TableName() string {
	return "room"
}
