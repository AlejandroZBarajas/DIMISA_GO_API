package clavesDomain

import (
	claveEntity "DIMISA/src/claves/clavesDomain/entity"
)

type ClaveInterface interface {
	SearchClave(s string) ([]*claveEntity.ClaveEntity, error)
}
