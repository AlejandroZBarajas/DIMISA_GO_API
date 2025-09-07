package camasDomain

import "DIMISA/src/camas/camasDomain/camaEntity"

type CamaInterface interface {
	CreateCama(cama *camaEntity.CamaEntity) error
	UpdateCama(cama *camaEntity.CamaEntity) error
	GetCamasByArea(areaid int32) ([]*camaEntity.CamaEntity, error)
	EnableCama(id int32) error
	DisableCama(id int32) error
	DeleteCama(id int32) error
}
