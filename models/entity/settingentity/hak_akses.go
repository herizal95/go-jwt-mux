package settingentity

type HakAkses struct {
	ID        uint   `gorm:"uniqueIndex;primary_key" json:"id"`
	NamaAkses string `json:"nama_akses"`
	Allowed   string `json:"allowed"`
}
