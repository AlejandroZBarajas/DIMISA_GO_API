package camasApp

import (
	"DIMISA/src/camas/camasDomain"
	"DIMISA/src/camas/camasDomain/camaEntity"
)

type GetFreeCamasByArea struct {
	Repo camasDomain.CamaInterface
}

func (uc *GetFreeCamasByArea) Execute(area_id int32) ([]*camaEntity.CamaEntity, error) {
	return uc.Repo.GetFreeCamasByArea(area_id)
}
