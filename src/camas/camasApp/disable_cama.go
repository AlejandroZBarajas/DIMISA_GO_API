package camasApp

import (
	"DIMISA/src/camas/camasDomain"
)

type DisableCama struct {
	Repo camasDomain.CamaInterface
}

func (uc *DisableCama) Execute(id int32) error {
	return uc.Repo.DisableCama(id)
}
