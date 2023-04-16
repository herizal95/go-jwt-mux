package routes

import (
	"github.com/gorilla/mux"
	usercontroller "github.com/herizal95/hisabia_api/controllers/usercontroller"
	"github.com/herizal95/hisabia_api/middleware"
)

func UserRoutes(r *mux.Router) {

	router := r.PathPrefix("/user").Subrouter()

	// middleware
	router.Use(middleware.Deserializer)

	router.HandleFunc("/", usercontroller.CreateUser).Methods("POST")
	router.HandleFunc("/profile", usercontroller.MyProfile).Methods("GET")
	router.HandleFunc("/{uid}", usercontroller.GetUserID).Methods("GET")
	router.HandleFunc("/{uid}", usercontroller.UpdateUser).Methods("PUT")
	router.HandleFunc("/{uid}", usercontroller.DeleteUser).Methods("DELETE")

}
