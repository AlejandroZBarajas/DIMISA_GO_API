package cendisApp

import (
	"DIMISA/src/cendis/cendisDomain"
	cendisEntity "DIMISA/src/cendis/cendisDomain/entity"
)

type UpdateCendisUseCase struct {
	Repo cendisDomain.CendisInterface
}

func (uc *UpdateCendisUseCase) Execute(cendis *cendisEntity.CendisEntity, areas []int32) error {
	return uc.Repo.UpdateCendis(cendis, areas)
}
