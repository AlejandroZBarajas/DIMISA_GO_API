package camasApp

import (
	"DIMISA/src/camas/camasDomain"
)

type DeleteCama struct {
	Repo camasDomain.CamaInterface
}

func (uc *DeleteCama) Execute(id int32) error {
	return uc.Repo.DeleteCama(id)
}
