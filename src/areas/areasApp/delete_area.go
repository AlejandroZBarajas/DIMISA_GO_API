package areasApp

import (
	"DIMISA/src/areas/areasDomain"
)

type DeleteAreaUseCase struct {
	Repo areasDomain.AreasInterface
}

func (uc *DeleteAreaUseCase) Execute(id int32) error {
	return uc.Repo.DeleteArea(id)
}
