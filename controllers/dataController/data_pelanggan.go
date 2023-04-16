package datacontroller

import (
	"encoding/json"
	"net/http"

	"github.com/herizal95/hisabia_api/config"
	"github.com/herizal95/hisabia_api/helper"
	"github.com/herizal95/hisabia_api/models/entity/dataentity"
)

func FindDataPelanggan(w http.ResponseWriter, r *http.Request) {

	var data []dataentity.DataPelanggan

	if err := config.DB.
		Preload("Harga").Preload("Outlet").Preload("Usaha").
		Joins("LEFT OUTER JOIN data_hargas ON data_pelanggans.id_harga = data_hargas.id").
		Joins("LEFT OUTER JOIN data_outlets ON data_pelanggans.id_outlet = data_outlets.id").
		Joins("LEFT OUTER JOIN data_usahas ON data_pelanggans.id_usaha = data_usahas.id").
		Find(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, "Data Pelanggan Not Found", nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully fetch all data Pelanggan", data)

}

func FindByIdPelanggan(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var data dataentity.DataPelanggan

	if err := config.DB.
		Preload("Usaha").Preload("Outlet").Preload("Harga").
		Joins("left join data_outlets on data_pelanggans.id_outlet = data_outlets.id").
		Joins("left join data_usahas on data_pelanggans.id_usaha = data_usahas.id").
		Joins("LEFT OUTER JOIN data_hargas ON data_pelanggans.id_harga = data_hargas.id").
		Where("data_pelanggans.id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusNotFound, "Data Pelanggan not found", nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully Get Data by ID", data)

}

func CreatePelanggan(w http.ResponseWriter, r *http.Request) {

	var input dataentity.DataPelanggan

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	data := dataentity.DataPelanggan{
		// ID:          uuid.New(),
		NamaPelanggan: input.NamaPelanggan,
		Alamat:        input.Alamat,
		Contact:       input.Contact,
		IDHarga:       input.IDHarga,
		IDOutlet:      input.IDOutlet,
		IDUsaha:       input.IDUsaha,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, 201, "Successfully to Create data Pelanggan", nil)

}

func UpdatePelanggan(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var data dataentity.DataPelanggan
	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var datas dataentity.DataPelanggan
	if err := json.NewDecoder(r.Body).Decode(&datas); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	data.NamaPelanggan = datas.NamaPelanggan
	data.Alamat = datas.Alamat
	data.Contact = datas.Contact
	data.IDHarga = datas.IDHarga
	data.IDUsaha = datas.IDUsaha
	data.IDOutlet = datas.IDOutlet

	if err := config.DB.Updates(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully to Update data", data)

}

func DeletePelanggan(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var data dataentity.DataPelanggan
	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := config.DB.Delete(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
	}

	helper.ResponseJson(w, http.StatusOK, "Successfull Delete Data Pelanggan", nil)
}
