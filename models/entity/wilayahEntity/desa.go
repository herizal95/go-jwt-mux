package wilayahEntity

type Desa struct {
	ID          int64     `json:"id" gorm:"primary_key"`
	Kecamatan   Kecamatan `json:"kecamatan" gorm:"foreignKey:IDKecamatan"`
	IDKecamatan int64     `json:"id_kecamatan"`
	NamaDesa    string    `json:"nama_desa"`
}
