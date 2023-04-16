package transaksientity

import (
	"time"

	"github.com/herizal95/hisabia_api/models/entity/dataentity"
)

type Transaksi struct {
	ID              uint                     `gorm:"uniqueIndex;primary_key" json:"id"`
	KodeTransaksi   string                   `json:"kode_transaksi"`
	Tanggal         time.Time                `json:"tanggal"`
	StatusTransaksi string                   `json:"status_transaksi"`
	TotalHarga      float32                  `json:"status_harga"`
	Terhuntang      float32                  `json:"terhuntang"`
	Pelanggan       dataentity.DataPelanggan `gorm:"foreignKey:IDPelanggan" json:"data_pelanggan,omitempty"`
	IDPelanggan     int                      `json:"id_pelanggan,omitempty"`
	Outlet          dataentity.DataOutlet    `gorm:"foreignKey:IDOutlet" json:"data_outlet,omitempty"`
	IDOutlet        int                      `json:"id_outlet,omitempty"`
	Usaha           dataentity.DataUsaha     `gorm:"foreignKey:IDUsaha" json:"data_usaha,omitempty"`
	IDUsaha         int                      `json:"id_usaha,omitempty"`
	Harga           dataentity.DataHarga     `gorm:"foreignKey:IDHarga" json:"data_harga,omitempty"`
	IDHarga         int                      `json:"id_harga,omitempty"`
}
