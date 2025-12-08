package colectivosApp

import (
	"DIMISA/src/colectivos/colectivosDomain"
	"DIMISA/src/colectivos/colectivosDomain/colectivoEntity"
)

type AddToColectivo struct {
	Repo colectivosDomain.ColectivoInterface
}

func (uc *AddToColectivo) Execute(colectivo int32, claves []*colectivoEntity.ColectivoDetalleEntity) error {
	return uc.Repo.AddToColectivo(colectivo, claves)
}
