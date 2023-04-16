package wilayahEntity

type Kabupaten struct {
	ID            int64    `json:"id" gorm:"primary_key"`
	Provinsi      Provinsi `json:"provinsi" gorm:"foreignKey:IDProvinsi"`
	IDProvinsi    int64    `json:"id_provinsi"`
	NamaKabupaten string   `json:"nama_kabupaten"`
}
