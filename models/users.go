package models

import (
	"github.com/google/uuid"
)

type User struct {
	Uuid        uuid.UUID `gorm:"primary_key" json:"uuid"`
	NamaLengkap string    `json:"nama_lengkap" gorm:"varchar(100)"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
}
