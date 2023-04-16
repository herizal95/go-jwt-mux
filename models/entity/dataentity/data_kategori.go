package dataentity

import "time"

type DataKategori struct {
	ID           uint       `gorm:"uniqueIndex;primary_key"`
	NamaKategori string     `json:"nama_kategori"`
	Outlet       DataOutlet `gorm:"foreignKey:IDOutlet" json:"data_outlet,omitempty"`
	IDOutlet     uint       `json:"id_outlet"`
	Usaha        DataUsaha  `gorm:"foreignKey:IDUsaha" json:"data_usaha,omitempty"`
	IDUsaha      uint       `json:"id_usaha"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
