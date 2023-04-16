package dataroutes

import (
	"github.com/gorilla/mux"
	datacontroller "github.com/herizal95/hisabia_api/controllers/dataController"
)

func DataRoutes(r *mux.Router) {

	router := r.PathPrefix("/data").Subrouter()

	// router.Use(middleware.Deserializer)

	// DATA USAHA
	router.HandleFunc("/usahas", datacontroller.FindUsahas).Methods("GET")
	router.HandleFunc("/usahas", datacontroller.CreateUsaha).Methods("POST")
	router.HandleFunc("/usahas/{id}", datacontroller.UpdateUsaha).Methods("PUT")
	router.HandleFunc("/usahas/{id}", datacontroller.FindByIdUsaha).Methods("GET")
	router.HandleFunc("/usahas/{id}", datacontroller.DeleteUsaha).Methods("DELETE")

	// DATA OUTLET
	router.HandleFunc("/outlets", datacontroller.FindOutlets).Methods("GET")
	router.HandleFunc("/outlets", datacontroller.CreatedOutlet).Methods("POST")
	router.HandleFunc("/outlets/{id}", datacontroller.UpdatedOutlet).Methods("PUT")
	router.HandleFunc("/outlets/{id}", datacontroller.FindByIdOutlet).Methods("GET")
	router.HandleFunc("/outlets/{id}", datacontroller.DeletedOutlet).Methods("DELETE")

	// DATA HARGA
	router.HandleFunc("/hargas", datacontroller.FindDataHargas).Methods("GET")
	router.HandleFunc("/hargas", datacontroller.CreatedHarga).Methods("POST")
	router.HandleFunc("/hargas/{id}", datacontroller.UpdatedHarga).Methods("PUT")
	router.HandleFunc("/hargas/{id}", datacontroller.FindByIdHarga).Methods("GET")
	router.HandleFunc("/hargas/{id}", datacontroller.DeletedHarga).Methods("DELETE")

	// DATA CATEGORY
	router.HandleFunc("/kategoris", datacontroller.FindCategories).Methods("GET")
	router.HandleFunc("/kategoris", datacontroller.CreateCategories).Methods("POST")
	router.HandleFunc("/kategoris/{id}", datacontroller.UpdateCategories).Methods("PUT")
	router.HandleFunc("/kategoris/{id}", datacontroller.FindByIdCategories).Methods("GET")
	router.HandleFunc("/kategoris/{id}", datacontroller.DeleteCategories).Methods("DELETE")

	// DATA SUPPLIER
	router.HandleFunc("/suppliers", datacontroller.FindSuppliers).Methods("GET")
	router.HandleFunc("/suppliers", datacontroller.CreateSupplier).Methods("POST")
	router.HandleFunc("/suppliers/{id}", datacontroller.UpdateSupplier).Methods("PUT")
	router.HandleFunc("/suppliers/{id}", datacontroller.FindByIdSupplier).Methods("GET")
	router.HandleFunc("/suppliers/{id}", datacontroller.DeleteSupplier).Methods("DELETE")

	// DATA SATUAN
	router.HandleFunc("/satuans", datacontroller.FindDataSatuan).Methods("GET")
	router.HandleFunc("/satuans", datacontroller.CreatedSatuan).Methods("POST")
	router.HandleFunc("/satuans/{id}", datacontroller.UpdatedSatuan).Methods("PUT")
	router.HandleFunc("/satuans/{id}", datacontroller.FindByIdSatuan).Methods("GET")
	router.HandleFunc("/satuans/{id}", datacontroller.DeletedSatuan).Methods("DELETE")

	// DATA PELANGGAN
	router.HandleFunc("/pelanggans", datacontroller.FindDataPelanggan).Methods("GET")
	router.HandleFunc("/pelanggans", datacontroller.CreatePelanggan).Methods("POST")
	router.HandleFunc("/pelanggans/{id}", datacontroller.UpdatePelanggan).Methods("PUT")
	router.HandleFunc("/pelanggans/{id}", datacontroller.FindByIdPelanggan).Methods("GET")
	router.HandleFunc("/pelanggans/{id}", datacontroller.DeletePelanggan).Methods("DELETE")
}
