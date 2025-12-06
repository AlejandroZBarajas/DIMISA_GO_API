package camasApp

import (
	"DIMISA/src/camas/camasDomain"
)

type SetFreeCama struct {
	Repo camasDomain.CamaInterface
}

func (uc *SetFreeCama) Execute(id_cama int32) error {
	return uc.Repo.SetFreeCama(id_cama)
}
