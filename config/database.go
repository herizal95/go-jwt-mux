package config

import (
	"fmt"
	"log"
	"os"

	"github.com/herizal95/hisabia_api/models/entity"
	"github.com/herizal95/hisabia_api/models/entity/dataentity"
	"github.com/herizal95/hisabia_api/models/entity/transaksientity"
	"github.com/herizal95/hisabia_api/models/entity/wilayahEntity"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("error Loading .env File")

	}
	// Mengambil nilai variabel lingkungan
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	// Koneksi ke database menggunakan GORM
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	// migrate models to database postgres

	DB.AutoMigrate(
		&entity.User{},
		&wilayahEntity.Provinsi{},
		&wilayahEntity.Kabupaten{},
		&wilayahEntity.Kecamatan{},
		&wilayahEntity.Desa{},
		&dataentity.DataUsaha{},
		&dataentity.DataOutlet{},
		&dataentity.DataHarga{},
		&dataentity.DataKategori{},
		&dataentity.DataSupplier{},
		&dataentity.DataPelanggan{},
		&dataentity.DataSatuan{},
		&transaksientity.Transaksi{},
	)

	if err != nil {
		return
	}
}
