package camasApp

import (
	"DIMISA/src/camas/camasDomain"
	"DIMISA/src/camas/camasDomain/camaEntity"
)

type UpdateCama struct {
	Repo camasDomain.CamaInterface
}

func (uc *UpdateCama) Execute(cama *camaEntity.CamaEntity) error {
	return uc.Repo.UpdateCama(cama)
}
