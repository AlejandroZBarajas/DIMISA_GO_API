package colectivosApp

import (
	"DIMISA/src/colectivos/colectivosDomain"
	"DIMISA/src/colectivos/colectivosDomain/colectivoEntity"
)

type CreateColectivo struct {
	Repo colectivosDomain.ColectivoInterface
}

func (uc *CreateColectivo) Execute(colectivo *colectivoEntity.ColectivoEntity) error {
	return uc.Repo.CreateColectivo(colectivo)
}
