package camasApp

import (
	"DIMISA/src/camas/camasDomain"
	"DIMISA/src/camas/camasDomain/camaEntity"
)

type CreateCamasRange struct {
	Repo camasDomain.CamaInterface
}

func (uc *CreateCamasRange) Execute(idArea, cama1, camaN int32) error {

	if camaN == 0 || camaN == cama1 {
		cama := &camaEntity.CamaEntity{
			Id_area:     idArea,
			Numero_cama: cama1,
		}
		return uc.Repo.CreateCama(cama)
	}

	for i := cama1; i <= camaN; i++ {
		cama := &camaEntity.CamaEntity{
			Id_area:     idArea,
			Numero_cama: i,
		}
		if err := uc.Repo.CreateCama(cama); err != nil {
			return err
		}
	}

	return nil
}
