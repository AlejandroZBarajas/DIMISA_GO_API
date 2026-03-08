package colectivosApp

import (
	"DIMISA/src/colectivos/colectivosDomain"
)

type CloseColectivo struct {
	Repo colectivosDomain.ColectivoInterface
}

func (uc *CloseColectivo) Execute(id int32) error {
	return uc.Repo.CloseColectivo(id)
}
