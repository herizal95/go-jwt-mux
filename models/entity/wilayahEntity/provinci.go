package wilayahEntity

type Provinsi struct {
	ID           int    `json:"id" gorm:"primary_key"`
	NamaProvinsi string `json:"nama_provinsi"`
}
