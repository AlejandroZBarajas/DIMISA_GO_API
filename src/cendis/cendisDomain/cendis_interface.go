package cendisDomain

import (
	cendisEntity "DIMISA/src/cendis/cendisDomain/entity"
)

type CendisInterface interface {
	CreateCendis(cendis *cendisEntity.CendisEntity, areas []int32) error
	UpdateCendis(cendis *cendisEntity.CendisEntity, areas []int32) error
	GetAllCendis() ([]*cendisEntity.CendisEntity, error)
	DeleteCendis(id int32) error
}
