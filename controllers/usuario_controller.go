package controllers

import (
	"API_THEHOUSEFIT/config"
	"API_THEHOUSEFIT/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Helper JSON
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func GetAllUsuarios(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id_usuario, nombres, apellidos, email, fecha_nacimiento, nacionalidad, ciudad, departamento, direccion, telefono, activo, fecha_modificacion, fecha_creacion FROM usuarios WHERE 1=1"

	nombres := r.URL.Query().Get("nombres")
	email := r.URL.Query().Get("email")

	if nombres != "" {
		query += " AND nombres ILIKE '%" + nombres + "%'"
	}
	if email != "" {
		query += " AND email ILIKE '%" + email + "%'"
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var usuarios []models.Usuario
	for rows.Next() {
		var u models.Usuario
		rows.Scan(
			&u.IDUsuario, &u.Nombres, &u.Apellidos, &u.Email,
			&u.FechaNacimiento, &u.Nacionalidad, &u.Ciudad,
			&u.Departamento, &u.Direccion, &u.Telefono,
			&u.Activo, &u.FechaModificacion, &u.FechaCreacion,
		)
		usuarios = append(usuarios, u)
	}
	respondJSON(w, 200, usuarios)
}
