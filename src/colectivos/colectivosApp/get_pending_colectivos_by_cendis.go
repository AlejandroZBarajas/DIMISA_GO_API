package colectivosApp

import (
	"DIMISA/src/colectivos/colectivosDomain"
	"DIMISA/src/colectivos/colectivosDomain/colectivoEntity"
)

type GetPendingColectivosByCendis struct {
	Repo colectivosDomain.ColectivoInterface
}

func (uc *GetPendingColectivosByCendis) Execute(id int32) ([]*colectivoEntity.ColectivoEntity, error) {
	return uc.Repo.GetPendingColectivosByCendis(id)
}
