package datacontroller

import (
	"encoding/json"
	"net/http"

	"github.com/herizal95/hisabia_api/config"
	"github.com/herizal95/hisabia_api/helper"
	"github.com/herizal95/hisabia_api/models/entity/dataentity"
)

func FindCategories(w http.ResponseWriter, r *http.Request) {
	var data []dataentity.DataKategori

	if err := config.DB.
		Preload("Outlet").Preload("Usaha").
		Joins("LEFT OUTER JOIN data_outlets ON data_kategoris.id_outlet = data_outlets.id").
		Joins("LEFT OUTER JOIN data_usahas ON data_kategoris.id_usaha = data_usahas.id").
		Find(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, "Data Kategori Not Found", nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully fetch all data Kategori", data)
}

func FindByIdCategories(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var data dataentity.DataKategori

	if err := config.DB.
		Preload("Usaha").Preload("Outlet").
		Joins("left join data_outlets on data_kategoris.id_outlet = data_outlets.id").
		Joins("left join data_usahas on data_kategoris.id_usaha = data_usahas.id").Where("data_kategoris.id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusNotFound, "Data Kategori not found", nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully Get Data by ID", data)

}

func CreateCategories(w http.ResponseWriter, r *http.Request) {

	var input dataentity.DataKategori

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	data := dataentity.DataKategori{
		// ID:          uuid.New(),
		NamaKategori: input.NamaKategori,
		IDOutlet:     input.IDOutlet,
		IDUsaha:      input.IDUsaha,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, 201, "Successfully to Create data Kategori", nil)
}

func UpdateCategories(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var data dataentity.DataKategori
	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var datas dataentity.DataKategori
	if err := json.NewDecoder(r.Body).Decode(&datas); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	data.NamaKategori = datas.NamaKategori
	data.IDUsaha = datas.IDUsaha
	data.IDOutlet = datas.IDOutlet

	if err := config.DB.Updates(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully to Update data", data)

}

func DeleteCategories(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var data dataentity.DataKategori
	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := config.DB.Delete(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
	}

	helper.ResponseJson(w, http.StatusOK, "Successfull Delete Data Kategori", nil)

}
