package clavesDomain

import (
	claveEntity "DIMISA/src/claves/clavesDomain/entity"
)

type ClaveInterface interface {
	SearchClave(s string) ([]*claveEntity.ClaveEntity, error)
	SearchInInventory(s string, id int32) ([]*claveEntity.ClaveEntity, error)
}
