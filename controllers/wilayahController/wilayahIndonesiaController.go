package wilayahConttroller

import (
	"net/http"

	"github.com/herizal95/hisabia_api/config"
	"github.com/herizal95/hisabia_api/helper"
	"github.com/herizal95/hisabia_api/models/entity/wilayahEntity"
)

func FindDesa(w http.ResponseWriter, r *http.Request) {

	var desa []wilayahEntity.Desa

	if err := config.DB.Model(&desa).Joins("LEFT JOIN kecamatans ON desas.id_kecamatan = kecamatans.id").
		Find(&desa).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	config.DB.Preload("Kecamatan").Find(&desa)

	helper.ResponseJson(w, http.StatusOK, "Successfully fetch all desa data", desa)
}

func FindKecamatan(w http.ResponseWriter, r *http.Request) {

	var kecamtan []wilayahEntity.Kecamatan

	if err := config.DB.
		Preload("Kabupaten").
		Joins("LEFT OUTER JOIN kabupatens ON kecamatans.id_kabupaten = kabupatens.id").
		Find(&kecamtan).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully fetch all kecamatan data", kecamtan)
}

func FindKabupaten(w http.ResponseWriter, r *http.Request) {

	var kabupaten []wilayahEntity.Kabupaten

	if err := config.DB.
		Preload("Provinsi").
		Joins("LEFT OUTER JOIN provinsis ON kabupatens.id_provinsi = provinsis.id").
		Find(&kabupaten).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully fetch all kabupaten data", kabupaten)
}

func FindProvinsi(w http.ResponseWriter, r *http.Request) {

	var provinsi []wilayahEntity.Provinsi

	if err := config.DB.Find(&provinsi).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully fetch all provinsi data", provinsi)
}
