package areasApp

import (
	"DIMISA/src/areas/areasDomain"
	"DIMISA/src/areas/areasDomain/areaEntity"
)

type GetAreasByCendisUseCase struct {
	Repo areasDomain.AreasInterface
}

func (uc *GetAreasByCendisUseCase) Execute(id int32) ([]*areaEntity.AreaEntity, error) {
	return uc.Repo.GetAreasByCendis(id)
}
