package colectivosDomain

import (
	colectivoEntity "DIMISA/src/colectivos/colectivosDomain/colectivoEntity"
)

type ColectivoInterface interface {
	CreateColectivo(colectivo *colectivoEntity.ColectivoEntity) error
	GetColectivosByCendis(id int32) ([]*colectivoEntity.ColectivoDTO, error)
	GetPendingColectivosByCendis(id int32) ([]*colectivoEntity.ColectivoDTO, error)
	GetUpdatableColectivosByCendis(id int32) ([]*colectivoEntity.ColectivoDTO, error)
}
