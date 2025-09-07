package camasApp

import (
	"DIMISA/src/camas/camasDomain"
	"DIMISA/src/camas/camasDomain/camaEntity"
)

type GetCamasByArea struct {
	Repo camasDomain.CamaInterface
}

func (uc *GetCamasByArea) Execute(id int32) ([]*camaEntity.CamaEntity, error) {
	return uc.Repo.GetCamasByArea(id)
}
