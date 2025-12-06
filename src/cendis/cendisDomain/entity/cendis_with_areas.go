package cendisEntity

import (
	"DIMISA/src/areas/areasDomain/areaEntity"
)

type CendisWithAreas struct {
	Id_cendis     int32                   `json:"id_cendis"`
	Cendis_nombre string                  `json:"cendis_nombre"`
	Areas         []areaEntity.AreaEntity `json:"areas"`
}
