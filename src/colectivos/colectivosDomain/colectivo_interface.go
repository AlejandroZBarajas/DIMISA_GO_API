package colectivosDomain

import (
	colectivoEntity "DIMISA/src/colectivos/colectivosDomain/colectivoEntity"
)

type ColectivoInterface interface {
	CreateColectivo(colectivo *colectivoEntity.ColectivoEntity) error
	GetColectivosByCendis(id int32) ([]*colectivoEntity.ColectivoDTO, error)
	GetPendingColectivosByCendis(id int32) ([]*colectivoEntity.ColectivoDTO, error)
	GetUpdatableColectivosByCendis(id int32) ([]*colectivoEntity.ColectivoDTO, error)
	AddToColectivo(id_cendis, tipo int32, claves []*colectivoEntity.ColectivoDetalleEntity) error
	CloseColectivo(id_colectivo int32) error
}
