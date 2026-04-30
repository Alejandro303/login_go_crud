package models

import "time"

type Usuario struct {
	IDUsuario        int        `json:"id_usuario"`
	Nombres          string     `json:"nombres"`
	Apellidos        string     `json:"apellidos"`
	Email            string     `json:"email"`
	FechaNacimiento  *time.Time `json:"fecha_nacimiento"`
	Nacionalidad     string     `json:"nacionalidad"`
	Ciudad           string     `json:"ciudad"`
	Departamento     string     `json:"departamento"`
	Direccion        string     `json:"direccion"`
	Telefono         string     `json:"telefono"`
	Activo           bool       `json:"activo"`
	FechaModificacion *time.Time `json:"fecha_modificacion"`
	FechaCreacion    *time.Time `json:"fecha_creacion"`
}