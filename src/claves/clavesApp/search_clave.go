package clavesApp

import (
	"DIMISA/src/claves/clavesDomain"
	claveEntity "DIMISA/src/claves/clavesDomain/entity"
)

type SearchClave struct {
	Repo clavesDomain.ClaveInterface
}

func (uc *SearchClave) Execute(s string) ([]*claveEntity.ClaveEntity, error) {
	return uc.Repo.SearchClave(s)
}
