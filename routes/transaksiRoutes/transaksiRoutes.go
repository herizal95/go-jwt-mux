package transaksiroutes

import (
	"github.com/gorilla/mux"
	transaksicontroller "github.com/herizal95/hisabia_api/controllers/transaksiController"
)

func TransaksiRoutes(r *mux.Router) {

	router := r.PathPrefix("/transaksis").Subrouter()

	// router.Use(middleware.Deserializer)

	router.HandleFunc("", transaksicontroller.FindTransaksis).Methods("GET")
	router.HandleFunc("", transaksicontroller.CreatedTransaksi).Methods("POST")
	router.HandleFunc("/pos", transaksicontroller.POSTransaksi).Methods("GET")
}
