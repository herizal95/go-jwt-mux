package dataentity

import "github.com/herizal95/hisabia_api/models/entity/settingentity"

type DataUser struct {
	ID         uint                   `gorm:"uniqueIndex;primary_key" json:"id"`
	Usernamer  string                 `json:"username"`
	Password   string                 `json:"password"`
	Outlet     DataOutlet             `gorm:"foreignKey:IDOutlet" json:"data_outlet,omitempty"`
	IDOutlet   int                    `json:"id_outlet"`
	HakAkses   settingentity.HakAkses `gorm:"foreignKey:IDHakAkses" json:"hak_akses,omitempty"`
	IDHakAkses int                    `json:"id_hak_akses"`
}
