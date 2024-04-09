package models

import (
	"github.com/eduhub/util"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string          `gorm:"type:varchar(255)"`
	LastName  string          `gorm:"type:varchar(255)"`
	UserName  string          `gorm:"type:varchar(255);unique_index"`
	Phone     string          `gorm:"type:varchar(255)"`
	Address   string          `gorm:"type:varchar(255)"`
	Password  string          `gorm:"type:varchar(255)"`
	DeletedAt *gorm.DeletedAt `gorm:"index;unique_index"`
	UserType  util.UserType
}

func (*User) TableName() string {
	return "users"
}
