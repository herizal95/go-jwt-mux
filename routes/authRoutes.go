package routes

import (
	"github.com/gorilla/mux"
	"github.com/herizal95/hisabia_api/controllers/auth"
)

func AuthenticationRoutes(r *mux.Router) {

	router := r.PathPrefix("/auth").Subrouter()

	router.HandleFunc("/register", auth.Register).Methods("POST")
	router.HandleFunc("/login", auth.Login).Methods("POST")
	// router.HandleFunc("/logout", auth.Logout).Methods("GET")
}
