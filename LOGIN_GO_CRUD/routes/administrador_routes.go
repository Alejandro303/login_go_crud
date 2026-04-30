package routes

import (
	"API_THEHOUSEFIT/controllers"
	"github.com/gorilla/mux"
)

func RegisterAdministradorRoutes(r *mux.Router) {
	r.HandleFunc("/administradores", controllers.GetAllAdministradores).Methods("GET")
	r.HandleFunc("/administradores/{id}", controllers.GetAdministradorByID).Methods("GET")
	r.HandleFunc("/administradores", controllers.CreateAdministrador).Methods("POST")
	r.HandleFunc("/administradores/{id}", controllers.UpdateAdministrador).Methods("PUT")
	r.HandleFunc("/administradores/{id}", controllers.DeleteAdministrador).Methods("DELETE")
}
