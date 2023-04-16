package dataentity

import "time"

type DataSupplier struct {
	ID           uint       `gorm:"uniqueIndex;primary_key"`
	NamaSupplier string     `json:"nama_supplier"`
	Alamat       string     `json:"alamat"`
	Contact      string     `json:"contact"`
	Outlet       DataOutlet `gorm:"foreignKey:IDOutlet" json:"data_outlet,omitempty"`
	IDOutlet     uint       `json:"id_outlet"`
	Usaha        DataUsaha  `gorm:"foreignKey:IDUsaha" json:"data_usaha,omitempty"`
	IDUsaha      uint       `json:"id_usaha"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
