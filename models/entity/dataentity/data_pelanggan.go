package dataentity

import "time"

type DataPelanggan struct {
	ID            uint       `gorm:"uniqueIndex;primary_key"`
	NamaPelanggan string     `json:"nama_pelanggan"`
	Alamat        string     `json:"alamat"`
	Contact       string     `json:"contact"`
	Harga         DataHarga  `gorm:"foreignKey:IDHarga" json:"data_harga,omitempty"`
	IDHarga       uint       `json:"id_harga"`
	Outlet        DataOutlet `gorm:"foreignKey:IDOutlet" json:"data_outlet,omitempty"`
	IDOutlet      uint       `json:"id_outlet"`
	Usaha         DataUsaha  `gorm:"foreignKey:IDUsaha" json:"data_usaha,omitempty"`
	IDUsaha       uint       `json:"id_usaha"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
