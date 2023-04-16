package datacontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/herizal95/hisabia_api/config"
	"github.com/herizal95/hisabia_api/helper"
	"github.com/herizal95/hisabia_api/models/entity/dataentity"
)

func FindUsahas(w http.ResponseWriter, r *http.Request) {

	var data []dataentity.DataUsaha

	if err := config.DB.Find(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully fetch all data Usaha", data)
}

func CreateUsaha(w http.ResponseWriter, r *http.Request) {

	var input dataentity.DataUsaha

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	dataUsaha := dataentity.DataUsaha{
		// ID:           uuid.New(),
		NamaUsaha:    input.NamaUsaha,
		Alamat:       input.Alamat,
		Contact:      input.Contact,
		EndTime:      time.Now(),
		Subscription: "full",
		IsActive:     1,
	}

	if err := config.DB.Create(&dataUsaha).Error; err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully to Create data usaha", dataUsaha)

}

func UpdateUsaha(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var data dataentity.DataUsaha
	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	var dataUpdate dataentity.DataUsaha
	if err := json.NewDecoder(r.Body).Decode(&dataUpdate); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	data.NamaUsaha = dataUpdate.NamaUsaha
	data.Alamat = dataUpdate.Alamat
	data.Contact = dataUpdate.Contact

	if err := config.DB.Updates(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully to Update User", data)
}

func FindByIdUsaha(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var data dataentity.DataUsaha

	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusNotFound, "Data Usaha ID not found", nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully Get Data by ID", data)
}

func DeleteUsaha(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var data dataentity.DataUsaha

	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusNotFound, "Data Usaha ID not found", nil)
		return
	}

	if err := config.DB.Delete(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfull Delete Data Usaha", nil)
}
