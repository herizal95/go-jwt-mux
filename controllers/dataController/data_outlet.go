package datacontroller

import (
	"encoding/json"
	"net/http"

	"github.com/herizal95/hisabia_api/config"
	"github.com/herizal95/hisabia_api/helper"
	"github.com/herizal95/hisabia_api/models/entity/dataentity"
)

func FindOutlets(w http.ResponseWriter, r *http.Request) {

	var data []dataentity.DataOutlet

	if err := config.DB.
		Preload("Usaha").Preload("Provinsi").Preload("Kabupaten").Preload("Kecamatan").Preload("Desa").
		Joins("LEFT OUTER JOIN data_usahas ON data_outlets.id_usaha = data_usahas.id").
		Joins("LEFT OUTER JOIN provinsis ON data_outlets.id_provinsi = provinsis.id").
		Joins("LEFT OUTER JOIN kabupatens ON data_outlets.id_kabupaten = kabupatens.id").
		Joins("LEFT OUTER JOIN kecamatans ON data_outlets.id_kecamatan = kecamatans.id").
		Joins("LEFT OUTER JOIN desas ON data_outlets.id_desa = desas.id").
		Find(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully fetch all data Outlet", data)
}

func FindByIdOutlet(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var data dataentity.DataOutlet

	if err := config.DB.
		Preload("Usaha").
		Joins("left join data_usahas on data_outlets.id_usaha = data_usahas.id").Where("data_outlets.id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusNotFound, "Data Outlet not found", nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully Get Data by ID", data)
}

func CreatedOutlet(w http.ResponseWriter, r *http.Request) {

	var input dataentity.DataOutlet

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	dataOutet := dataentity.DataOutlet{
		// ID:          uuid.New(),
		NamaOutlet:  input.NamaOutlet,
		Alamat:      input.Alamat,
		Contact:     input.Contact,
		Pic:         input.Pic,
		IDUsaha:     input.IDUsaha,
		IDProvinsi:  input.IDProvinsi,
		IDKabupaten: input.IDKabupaten,
		IDKecamatan: input.IDKecamatan,
		IDDesa:      input.IDDesa,
		IsPusat:     input.IsPusat,
	}

	if err := config.DB.Create(&dataOutet).Error; err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, 201, "Successfully to Create data outlet", nil)

}

func UpdatedOutlet(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var data dataentity.DataOutlet
	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var outlet dataentity.DataOutlet
	if err := json.NewDecoder(r.Body).Decode(&outlet); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	data.NamaOutlet = outlet.NamaOutlet
	data.Alamat = outlet.Alamat
	data.Contact = outlet.Contact
	data.Pic = outlet.Pic
	data.IDUsaha = outlet.IDUsaha
	data.IDProvinsi = outlet.IDProvinsi
	data.IDKabupaten = outlet.IDKabupaten
	data.IDKecamatan = outlet.IDKecamatan
	data.IDDesa = outlet.IDDesa

	if err := config.DB.Updates(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully to Update User", data)
}

func DeletedOutlet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var data dataentity.DataOutlet
	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := config.DB.Delete(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
	}

	helper.ResponseJson(w, http.StatusOK, "Successfull Delete Data Outlet", data)
}
