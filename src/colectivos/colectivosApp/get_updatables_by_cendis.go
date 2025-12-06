package colectivosApp

import (
	"DIMISA/src/colectivos/colectivosDomain"
	"DIMISA/src/colectivos/colectivosDomain/colectivoEntity"
)

type GetUpdatableColectivosByCendis struct {
	Repo colectivosDomain.ColectivoInterface
}

func (uc *GetUpdatableColectivosByCendis) Execute(id int32) ([]*colectivoEntity.ColectivoEntity, error) {
	return uc.Repo.GetUpdatableColectivosByCendis(id)
}
