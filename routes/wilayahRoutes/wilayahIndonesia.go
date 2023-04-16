package wilayahroutes

import (
	"github.com/gorilla/mux"
	wilayahConttroller "github.com/herizal95/hisabia_api/controllers/wilayahController"
)

func DesaRoutes(r *mux.Router) {

	router := r.PathPrefix("/wilayah-indonesia").Subrouter()

	router.HandleFunc("/desas", wilayahConttroller.FindDesa).Methods("GET")
	router.HandleFunc("/kecamatans", wilayahConttroller.FindKecamatan).Methods("GET")
	router.HandleFunc("/kabupatens", wilayahConttroller.FindKabupaten).Methods("GET")
	router.HandleFunc("/provinsis", wilayahConttroller.FindProvinsi).Methods("GET")
}
