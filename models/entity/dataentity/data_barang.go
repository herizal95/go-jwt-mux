package dataentity

type DataBarang struct {
	ID         uint         `gorm:"uniqueIndex;primary_key" json:"id"`
	NamaBarang string       `gorm:"type:varchar(100)" json:"nama_barang"`
	HargaBeli  float32      `json:"harga_beli"`
	HargaJual  float32      `json:"harga_jual"`
	Hpp        float32      `json:"hpp"`
	Satuan     DataSatuan   `gorm:"foreignKey:IDSatuan" json:"data_satuan,omitempty"`
	IDSatuan   int          `json:"id_satuan"`
	Supplier   DataSupplier `gorm:"foreignKey:IDSupplier" json:"data_supplier,omitempty"`
	IDSupplier int          `json:"id_supplier"`
	Kategori   DataKategori `gorm:"foreignKey:IDKategori" json:"data_kategori,omitempty"`
	IDKategori int          `json:"id_kategori"`
	Outlet     DataOutlet   `gorm:"foreignKey:IDOutlet" json:"data_outlet,omitempty"`
	IDOutlet   int          `json:"id_outlet"`
	Usaha      DataUsaha    `gorm:"foreignKey:IDUsaha" json:"data_usaha,omitempty"`
	IDUsaha    int          `json:"id_usaha"`
	Status     int          `json:"status"`
}
