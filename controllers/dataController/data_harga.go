package datacontroller

import (
	"encoding/json"
	"net/http"

	"github.com/herizal95/hisabia_api/config"
	"github.com/herizal95/hisabia_api/helper"
	"github.com/herizal95/hisabia_api/models/entity/dataentity"
)

func FindDataHargas(w http.ResponseWriter, r *http.Request) {

	var data []dataentity.DataHarga

	if err := config.DB.
		Preload("Outlet").Preload("Usaha").
		Joins("LEFT OUTER JOIN data_outlets ON data_hargas.id_outlet = data_outlets.id").
		Joins("LEFT OUTER JOIN data_usahas ON data_hargas.id_usaha = data_usahas.id").
		Find(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, "Data Harga Not Found", nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully fetch all data harga", data)

}

func FindByIdHarga(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var data dataentity.DataHarga

	if err := config.DB.
		Preload("Usaha").Preload("Outlet").
		Joins("left join data_outlets on data_hargas.id_outlet = data_outlets.id").
		Joins("left join data_usahas on data_hargas.id_usaha = data_usahas.id").Where("data_hargas.id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusNotFound, "Data Hargas not found", nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully Get Data by ID", data)
}

func CreatedHarga(w http.ResponseWriter, r *http.Request) {

	var input dataentity.DataHarga

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	data := dataentity.DataHarga{
		// ID:          uuid.New(),
		NamaHarga: input.NamaHarga,
		IDOutlet:  input.IDOutlet,
		IDUsaha:   input.IDUsaha,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, 201, "Successfully to Create data harga", nil)

}

func UpdatedHarga(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var data dataentity.DataHarga
	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var outlet dataentity.DataHarga
	if err := json.NewDecoder(r.Body).Decode(&outlet); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	data.NamaHarga = outlet.NamaHarga
	data.IDUsaha = outlet.IDUsaha
	data.IDOutlet = outlet.IDOutlet

	if err := config.DB.Updates(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully to Update data", data)
}

func DeletedHarga(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var data dataentity.DataHarga
	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := config.DB.Delete(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
	}

	helper.ResponseJson(w, http.StatusOK, "Successfull Delete Data Hargas", nil)
}
