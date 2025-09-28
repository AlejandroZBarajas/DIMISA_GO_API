package areasApp

import (
	"DIMISA/src/areas/areasDomain"
	"DIMISA/src/areas/areasDomain/areaEntity"
)

type GetFreeAreasUseCase struct {
	Repo areasDomain.AreasInterface
}

func (uc *GetFreeAreasUseCase) Execute() ([]*areaEntity.AreaEntity, error) {
	return uc.Repo.GetFreeAreas()
}
