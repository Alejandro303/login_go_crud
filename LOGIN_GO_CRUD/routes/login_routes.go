package routes

import (
	"API_THEHOUSEFIT/controllers"
	"github.com/gorilla/mux"
)

func RegisterLoginRoutes(r *mux.Router) {
	r.HandleFunc("/login", controllers.GetAllLogins).Methods("GET")
	r.HandleFunc("/login/{id}", controllers.GetLoginByID).Methods("GET")
	r.HandleFunc("/login", controllers.CreateLogin).Methods("POST")
	r.HandleFunc("/login/{id}", controllers.UpdateLogin).Methods("PUT")
	r.HandleFunc("/login/{id}", controllers.DeleteLogin).Methods("DELETE")
}
