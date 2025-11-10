package colectivosApp

import (
	"DIMISA/src/colectivos/colectivosDomain"
	"DIMISA/src/colectivos/colectivosDomain/colectivoEntity"
)

type GetColectivosByCendis struct {
	Repo colectivosDomain.ColectivoInterface
}

func (uc *GetColectivosByCendis) Execute(id int32) ([]*colectivoEntity.ColectivoEntity, error) {
	return uc.Repo.GetColectivosByCendis(id)
}
