package areasApp

import (
	"DIMISA/src/areas/areasDomain"
	"DIMISA/src/areas/areasDomain/areaEntity"
	"DIMISA/src/camas/camasDomain"
	"DIMISA/src/camas/camasDomain/camaEntity"
)

type CreateAreaUseCase struct {
	Repo     areasDomain.AreasInterface
	CamaRepo camasDomain.CamaInterface
}

func (uc *CreateAreaUseCase) Execute(area *areaEntity.AreaEntity) error {

	err := uc.Repo.CreateArea(area)
	if err != nil {
		return err
	}

	for i := area.Cama_1; i <= area.Cama_n; i++ {
		cama := &camaEntity.CamaEntity{
			Numero_cama: i,
			Id_area:     area.Id_area,
		}
		if err := uc.CamaRepo.CreateCama(cama); err != nil {
			return err
		}
	}

	return nil
}
