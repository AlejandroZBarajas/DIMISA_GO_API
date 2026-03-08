package clavesApp

import (
	"DIMISA/src/claves/clavesDomain"
	claveEntity "DIMISA/src/claves/clavesDomain/entity"
)

type SearchInInventory struct {
	Repo clavesDomain.ClaveInterface
}

func (uc *SearchInInventory) Execute(s string, id int32) ([]*claveEntity.ClaveEntity, error) {
	return uc.Repo.SearchInInventory(s, id)
}
