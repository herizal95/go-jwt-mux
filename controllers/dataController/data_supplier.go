package datacontroller

import (
	"encoding/json"
	"net/http"

	"github.com/herizal95/hisabia_api/config"
	"github.com/herizal95/hisabia_api/helper"
	"github.com/herizal95/hisabia_api/models/entity/dataentity"
)

func FindSuppliers(w http.ResponseWriter, r *http.Request) {

	var data []dataentity.DataSupplier

	if err := config.DB.
		Preload("Outlet").Preload("Usaha").
		Joins("LEFT OUTER JOIN data_outlets ON data_suppliers.id_outlet = data_outlets.id").
		Joins("LEFT OUTER JOIN data_usahas ON data_suppliers.id_usaha = data_usahas.id").
		Find(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, "Data Suppliers Not Found", nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully fetch all data Suppliers", data)

}

func FindByIdSupplier(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var data dataentity.DataSupplier

	if err := config.DB.
		Preload("Usaha").Preload("Outlet").
		Joins("left join data_outlets on data_suppliers.id_outlet = data_outlets.id").
		Joins("left join data_usahas on data_suppliers.id_usaha = data_usahas.id").Where("data_suppliers.id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusNotFound, "Data Suppliers not found", nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully Get Data by ID", data)

}

func CreateSupplier(w http.ResponseWriter, r *http.Request) {

	var input dataentity.DataSupplier

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	data := dataentity.DataSupplier{
		// ID:          uuid.New(),
		NamaSupplier: input.NamaSupplier,
		Alamat:       input.Alamat,
		Contact:      input.Contact,
		IDOutlet:     input.IDOutlet,
		IDUsaha:      input.IDUsaha,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, 201, "Successfully to Create data Suppliers", nil)

}

func UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var data dataentity.DataSupplier
	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var datas dataentity.DataSupplier
	if err := json.NewDecoder(r.Body).Decode(&datas); err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	data.NamaSupplier = datas.NamaSupplier
	data.Alamat = datas.Alamat
	data.Contact = datas.Contact
	data.IDUsaha = datas.IDUsaha
	data.IDOutlet = datas.IDOutlet

	if err := config.DB.Updates(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Successfully to Update data", data)

}

func DeleteSupplier(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var data dataentity.DataSupplier
	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := config.DB.Delete(&data).Error; err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, err.Error(), nil)
	}

	helper.ResponseJson(w, http.StatusOK, "Successfull Delete Data Suppliers", nil)
}
