package entity

import (
	"time"
)

type User struct {
	Uid       uint   `gorm:"uniqueIndex;primary_key" json:"uid"`
	Name      string `json:"name" gorm:"type:varchar(50)"`
	Email     string `json:"email" gorm:"uniqueIndex;not null"`
	Username  string `json:"username" gorm:"not null"`
	Passowrd  string `json:"password" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
