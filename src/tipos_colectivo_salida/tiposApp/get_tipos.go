package tiposApp

import (
	"DIMISA/src/tipos_colectivo_salida/tiposDomain"
	tipoEntity "DIMISA/src/tipos_colectivo_salida/tiposDomain/entity"
)

type GetTipos struct {
	Repo tiposDomain.TiposInterface
}

func (uc *GetTipos) Execute() (*[]tipoEntity.TipoEntity, error) {
	return uc.Repo.GetTipos()
}
