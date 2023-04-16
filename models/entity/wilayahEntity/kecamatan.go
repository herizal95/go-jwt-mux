package wilayahEntity

type Kecamatan struct {
	ID            int64     `json:"id" gorm:"primary_key"`
	Kabupaten     Kabupaten `json:"kabupaten" gorm:"foreignKey:IDKabupaten"`
	IDKabupaten   int64     `json:"id_kabupaten"`
	NamaKecamatan string    `json:"nama_kecamatan"`
}
