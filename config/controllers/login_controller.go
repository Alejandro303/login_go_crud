package controllers

import (
	"API_THEHOUSEFIT/config"
	"API_THEHOUSEFIT/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL logins
func GetAllLogins(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(
		"SELECT id_login, id_nombre, rol, activo, fecha_modificacion, fecha_creacion FROM login",
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Login
	for rows.Next() {
		var l models.Login
		rows.Scan(&l.IDLogin, &l.IDNombre, &l.Rol, &l.Activo, &l.FechaModificacion, &l.FechaCreacion)
		list = append(list, l)
	}
	respondJSON(w, 200, list)
}
