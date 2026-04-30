package controllers

import (
	"LOGIN_GO_CRUD/config"
	"LOGIN_GO_CRUD/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL administradores (con filtros opcionales)
func GetAllAdministradores(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_administrador, nombre_gym, nit, anio_fundacion, cantidad_clientes,
		departamento, ciudad, direccion, propietario_nombre, telefono, correo,
		pagina_web, firma, fecha_firma, estado, activo, fecha_modificacion, fecha_creacion
		FROM administrador WHERE 1=1`

	nombreGym := r.URL.Query().Get("nombre_gym")
	ciudad := r.URL.Query().Get("ciudad")

	if nombreGym != "" {
		query += " AND nombre_gym ILIKE '%" + nombreGym + "%'"
	}
	if ciudad != "" {
		query += " AND ciudad ILIKE '%" + ciudad + "%'"
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var admins []models.Administrador
	for rows.Next() {
		var a models.Administrador
		rows.Scan(
			&a.IDAdministrador, &a.NombreGym, &a.Nit, &a.AnioFundacion,
			&a.CantidadClientes, &a.Departamento, &a.Ciudad, &a.Direccion,
			&a.PropietarioNombre, &a.Telefono, &a.Correo, &a.PaginaWeb,
			&a.Firma, &a.FechaFirma, &a.Estado, &a.Activo,
			&a.FechaModificacion, &a.FechaCreacion,
		)
		admins = append(admins, a)
	}
	respondJSON(w, 200, admins)
}

func GetAdministradorByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var a models.Administrador

	err := config.DB.QueryRow(
		`SELECT id_administrador, nombre_gym, nit, anio_fundacion, cantidad_clientes,
		departamento, ciudad, direccion, propietario_nombre, telefono, correo,
		pagina_web, firma, fecha_firma, estado, activo, fecha_modificacion, fecha_creacion
		FROM administrador WHERE id_administrador=$1`,
		id,
	).Scan(
		&a.IDAdministrador, &a.NombreGym, &a.Nit, &a.AnioFundacion,
		&a.CantidadClientes, &a.Departamento, &a.Ciudad, &a.Direccion,
		&a.PropietarioNombre, &a.Telefono, &a.Correo, &a.PaginaWeb,
		&a.Firma, &a.FechaFirma, &a.Estado, &a.Activo,
		&a.FechaModificacion, &a.FechaCreacion,
	)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Administrador no encontrado"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, a)
}

func CreateAdministrador(w http.ResponseWriter, r *http.Request) {
	var a models.Administrador
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err := config.DB.QueryRow(
		`INSERT INTO administrador (nombre_gym, nit, anio_fundacion, cantidad_clientes,
		departamento, ciudad, direccion, propietario_nombre, telefono, correo,
		pagina_web, firma, fecha_firma, estado, activo)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) RETURNING id_administrador`,
		a.NombreGym, a.Nit, a.AnioFundacion, a.CantidadClientes,
		a.Departamento, a.Ciudad, a.Direccion, a.PropietarioNombre,
		a.Telefono, a.Correo, a.PaginaWeb, a.Firma,
		a.FechaFirma, a.Estado, a.Activo,
	).Scan(&a.IDAdministrador)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, a)
}

func UpdateAdministrador(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var a models.Administrador
	json.NewDecoder(r.Body).Decode(&a)

	_, err := config.DB.Exec(
		`UPDATE administrador SET nombre_gym=$1, nit=$2, anio_fundacion=$3, cantidad_clientes=$4,
		departamento=$5, ciudad=$6, direccion=$7, propietario_nombre=$8, telefono=$9, correo=$10,
		pagina_web=$11, firma=$12, fecha_firma=$13, estado=$14, activo=$15
		WHERE id_administrador=$16`,
		a.NombreGym, a.Nit, a.AnioFundacion, a.CantidadClientes,
		a.Departamento, a.Ciudad, a.Direccion, a.PropietarioNombre,
		a.Telefono, a.Correo, a.PaginaWeb, a.Firma,
		a.FechaFirma, a.Estado, a.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Administrador actualizado correctamente"})
}

func DeleteAdministrador(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM administrador WHERE id_administrador=$1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Administrador eliminado correctamente"})
}