package areasApp

import (
	"DIMISA/src/areas/areasDomain"
	"DIMISA/src/areas/areasDomain/areaEntity"
)

type GetAreaByIDUseCase struct {
	Repo areasDomain.AreasInterface
}

func (uc *GetAreaByIDUseCase) Execute(id int32) (*areaEntity.AreaEntity, error) {
	return uc.Repo.GetAreaByID(id)
}
