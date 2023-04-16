package dataentity

import (
	"time"

	"github.com/herizal95/hisabia_api/models/entity/wilayahEntity"
)

type DataOutlet struct {
	ID          uint                    `gorm:"size:36;not null;uniqueIndex;primary_key"`
	NamaOutlet  string                  `gorm:"size:100;not null" json:"nama_outlet"`
	Alamat      string                  `gorm:"size:255" json:"alamat"`
	Contact     string                  `json:"contact"`
	Pic         string                  `json:"pic"`
	Usaha       DataUsaha               `gorm:"foreignKey:IDUsaha" json:"usaha,omitempty"`
	IDUsaha     uint                    `json:"id_usaha,omitempty"`
	Provinsi    wilayahEntity.Provinsi  `gorm:"foreignKey:IDProvinsi" json:"provinsi,omitempty"`
	IDProvinsi  uint                    `json:"id_provinsi"`
	Kabupaten   wilayahEntity.Kabupaten `gorm:"foreignKey:IDKabupaten" json:"kabupaten,omitempty"`
	IDKabupaten uint                    `json:"id_kabupaten"`
	Kecamatan   wilayahEntity.Kecamatan `gorm:"foreignKey:IDKecamatan" json:"kecamatan,omitempty"`
	IDKecamatan uint                    `json:"id_kecamatan"`
	Desa        wilayahEntity.Desa      `gorm:"foreignKey:IDDesa" json:"desa,omitempty"`
	IDDesa      uint                    `json:"id_desa"`
	IsPusat     int                     `json:"is_pusat"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
