package colectivosDomain

import (
	colectivoEntity "DIMISA/src/colectivos/colectivosDomain/colectivoEntity"
)

type ColectivoInterface interface {
	CreateColectivo(colectivo *colectivoEntity.ColectivoEntity) error
	GetColectivosByCendis(id int32) ([]*colectivoEntity.ColectivoEntity, error)
	GetPendingColectivosByCendis(id int32) ([]*colectivoEntity.ColectivoEntity, error)
	GetUpdatableColectivosByCendis(id int32) ([]*colectivoEntity.ColectivoEntity, error)
}
