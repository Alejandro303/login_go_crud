package models

import "time"

type Administrador struct {
	IDAdministrador   int        `json:"id_administrador"`
	NombreGym         string     `json:"nombre_gym"`
	Nit               string     `json:"nit"`
	AnioFundacion     int        `json:"anio_fundacion"`
	CantidadClientes  int        `json:"cantidad_clientes"`
	Departamento      string     `json:"departamento"`
	Ciudad            string     `json:"ciudad"`
	Direccion         string     `json:"direccion"`
	PropietarioNombre string     `json:"propietario_nombre"`
	Telefono          string     `json:"telefono"`
	Correo            string     `json:"correo"`
	PaginaWeb         string     `json:"pagina_web"`
	Firma             string     `json:"firma"`
	FechaFirma        *time.Time `json:"fecha_firma"`
	Estado            string     `json:"estado"`
	Activo            bool       `json:"activo"`
	FechaModificacion *time.Time `json:"fecha_modificacion"`
	FechaCreacion     *time.Time `json:"fecha_creacion"`
}
