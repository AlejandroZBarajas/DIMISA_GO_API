package cendisApp

import (
	"DIMISA/src/cendis/cendisDomain"
	cendisEntity "DIMISA/src/cendis/cendisDomain/entity"
)

type CreateCendisUseCase struct {
	Repo cendisDomain.CendisInterface
}

func (uc *CreateCendisUseCase) Execute(cendis *cendisEntity.CendisEntity, areas []int32) error {
	return uc.Repo.CreateCendis(cendis, areas)
}
