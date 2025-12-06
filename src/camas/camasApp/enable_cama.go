package camasApp

import (
	"DIMISA/src/camas/camasDomain"
)

type EnableCama struct {
	Repo camasDomain.CamaInterface
}

func (uc *EnableCama) Execute(id int32) error {
	return uc.Repo.EnableCama(id)
}
