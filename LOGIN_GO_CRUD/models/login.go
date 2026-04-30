package models

import "time"

type Login struct {
	IDLogin           int        `json:"id_login"`
	IDNombre          int        `json:"id_nombre"`
	Rol               string     `json:"rol"`
	Activo            bool       `json:"activo"`
	FechaModificacion *time.Time `json:"fecha_modificacion"`
	FechaCreacion     *time.Time `json:"fecha_creacion"`
}
