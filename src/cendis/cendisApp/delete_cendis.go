package cendisApp

import (
	"DIMISA/src/cendis/cendisDomain"
)

type DeleteCendisUseCase struct {
	Repo cendisDomain.CendisInterface
}

func (uc *DeleteCendisUseCase) Execute(id int32) error {
	return uc.Repo.DeleteCendis(id)
}
