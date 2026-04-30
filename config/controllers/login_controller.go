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


func GetLoginByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var l models.Login

	err := config.DB.QueryRow(
		"SELECT id_login, id_nombre, rol, activo, fecha_modificacion, fecha_creacion FROM login WHERE id_login=$1",
		id,
	).Scan(&l.IDLogin, &l.IDNombre, &l.Rol, &l.Activo, &l.FechaModificacion, &l.FechaCreacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Login no encontrado"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, l)
}