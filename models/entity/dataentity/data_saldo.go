package dataentity

type DataSaldo struct {
	ID        uint       `gorm:"uniqueIndex;primary_key" json:"id"`
	NamaSaldo string     `json:"nama_saldo"`
	Keluar    float32    `json:"keluar"`
	Masuk     float32    `json:"masuk"`
	Outlet    DataOutlet `gorm:"foreignKey:IDOutlet" json:"data_outlet,omitempty"`
	IDOutlet  int        `json:"id_outlet"`
}
