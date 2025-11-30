package salidasApp

import (
	"DIMISA/src/salidas/salidasDomain"
	salidaEntity "DIMISA/src/salidas/salidasDomain/entity"
)

type GetSalidasByCendis struct {
	Repo salidasDomain.SalidasInterface
}

func (uc *GetSalidasByCendis) Execute(id_cendis int32) (*[]salidaEntity.SalidaEntity, error) {
	return uc.Repo.GetSalidasByCendis(id_cendis)
}
