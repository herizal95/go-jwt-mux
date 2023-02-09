package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("alphacsoft:alphacsoft123@tcp(localhost:3306)/golang_mux_starter"))
	if err != nil {
		fmt.Println("Gagal Koneksi database")
	}

	db.AutoMigrate(&User{})

	DB = db
}
