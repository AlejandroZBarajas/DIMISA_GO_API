package salidasApp

import (
	"DIMISA/src/salidas/salidasDomain"
	salidaEntity "DIMISA/src/salidas/salidasDomain/entity"
	"fmt"
)

type GetSalidasByCendis struct {
	Repo salidasDomain.SalidasInterface
}

func (uc *GetSalidasByCendis) Execute(id_cendis int32) (*[]salidaEntity.SalidaEntity, error) {
	fmt.Println("id_cendis: ", id_cendis)
	return uc.Repo.GetSalidasByCendis(id_cendis)
}
