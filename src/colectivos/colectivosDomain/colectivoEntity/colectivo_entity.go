package colectivoEntity

import "database/sql"

type ColectivoEntity struct {
	Tipo_id   int32                    `json:"tipo_id"`
	Fecha     string                   `json:"fecha"`
	Id_user   int32                    `json:"id_user"`
	Id_area   sql.NullInt32            `json:"id_area"`
	Id_cendis int32                    `json:"id_cendis"`
	Claves    []ColectivoDetalleEntity `json:"claves"`
}
