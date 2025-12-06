package cendisApp

import (
	"DIMISA/src/cendis/cendisDomain"
	cendisEntity "DIMISA/src/cendis/cendisDomain/entity"
)

type GetAllCendisUseCase struct {
	Repo cendisDomain.CendisInterface
}

func (uc *GetAllCendisUseCase) Execute() ([]*cendisEntity.CendisWithAreas, error) {
	return uc.Repo.GetAllCendis()
}
