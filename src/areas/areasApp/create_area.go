package areasApp

import (
	"DIMISA/src/areas/areasDomain"
	"DIMISA/src/areas/areasDomain/areaEntity"
	"DIMISA/src/camas/camasDomain"
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

	return nil
}
