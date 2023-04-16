package dataentity

type DataStok struct {
	ID       uint       `gorm:"uniqueIndex;primary_key" json:"id"`
	Outlet   DataOutlet `gorm:"foreignKey:IDOutlet" json:"data_outlet,omitempty"`
	IDOutlet int        `json:"id_outlet"`
	Barang   DataBarang `gorm:"foreignKey:IDBarang" json:"data_barang,omitempty"`
	IDBarang int        `json:"id_barang"`
	Stok     int        `json:"stok"`
	Usaha    DataUsaha  `gorm:"foreignKey:IDUsaha" json:"data_usaha,omitempty"`
	IDUsaha  int        `json:"id_usaha"`
}
