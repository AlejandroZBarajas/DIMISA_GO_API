package tiposDomain

import (
	tipoEntity "DIMISA/src/tipos_colectivo_salida/tiposDomain/entity"
)

type TiposInterface interface {
	GetTipos() (*[]tipoEntity.TipoEntity, error)
}
