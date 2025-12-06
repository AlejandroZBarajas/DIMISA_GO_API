package camasApp

import (
	"DIMISA/src/camas/camasDomain"
	"DIMISA/src/camas/camasDomain/camaEntity"
)

type CreateCama struct {
	Repo camasDomain.CamaInterface
}

func (uc *CreateCama) Execute(cama *camaEntity.CamaEntity) error {
	return uc.Repo.CreateCama(cama)
}
