package areasApp

import (
	"DIMISA/src/areas/areasDomain"
	"DIMISA/src/areas/areasDomain/areaEntity"
)

type UpdateAreaUseCase struct {
	Repo areasDomain.AreasInterface
}

func (uc *UpdateAreaUseCase) Execute(area *areaEntity.AreaEntity) error {
	return uc.Repo.UpdateArea(area)
}
