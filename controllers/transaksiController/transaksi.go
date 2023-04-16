package transaksicontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/herizal95/hisabia_api/config"
	"github.com/herizal95/hisabia_api/helper"
	"github.com/herizal95/hisabia_api/models/entity/transaksientity"
	"github.com/herizal95/hisabia_api/utils"
)

func FindTransaksis(w http.ResponseWriter, r *http.Request) {

	var dataTransaksi []transaksientity.Transaksi

	if err := config.DB.
		Preload("Pelanggan").Preload("Outlet").Preload("Harga").
		Joins("LEFT OUTER JOIN data_pelanggans ON transaksis.id_pelanggan = data_pelanggans.id").
		Joins("LEFT OUTER JOIN data_outlets ON transaksis.id_outlet = data_outlets.id").
		Joins("LEFT OUTER JOIN data_hargas ON transaksis.id_harga = data_hargas.id").
		Find(&dataTransaksi).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully fetch data transaksi", dataTransaksi)
}

func CreatedTransaksi(w http.ResponseWriter, r *http.Request) {

	var inputData transaksientity.Transaksi

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&inputData); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	kodeTransaksi := utils.GenerateKode()

	data := transaksientity.Transaksi{
		KodeTransaksi:   kodeTransaksi,
		Tanggal:         time.Now(),
		IDPelanggan:     inputData.IDPelanggan,
		IDOutlet:        inputData.IDOutlet,
		IDHarga:         inputData.IDHarga,
		IDUsaha:         inputData.IDUsaha,
		TotalHarga:      0,
		Terhuntang:      0,
		StatusTransaksi: "pending",
	}

	if err := config.DB.Create(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully to create transaksi", nil)
}

func POSTransaksi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POS Bayar")
}
