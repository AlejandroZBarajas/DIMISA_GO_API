package areasApp

import (
	"DIMISA/src/areas/areasDomain"
	"DIMISA/src/areas/areasDomain/areaEntity"
)

type GetAllAreasUseCase struct {
	Repo areasDomain.AreasInterface
}

func (uc *GetAllAreasUseCase) Execute() ([]*areaEntity.AreaEntity, error) {
	return uc.Repo.GetAllAreas()
}
