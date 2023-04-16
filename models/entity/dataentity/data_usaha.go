package dataentity

import (
	"time"
)

type DataUsaha struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	NamaUsaha    string    `json:"nama_usaha"`
	Alamat       string    `json:"alamat"`
	Contact      string    `json:"contact"`
	EndTime      time.Time `json:"endtime"`
	Subscription string    `json:"subscription"`
	IsActive     int       `json:"is_active"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
