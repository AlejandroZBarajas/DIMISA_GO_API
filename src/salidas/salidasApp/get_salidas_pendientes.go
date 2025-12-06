package salidasApp

import (
	"DIMISA/src/salidas/salidasDomain"
	salidaEntity "DIMISA/src/salidas/salidasDomain/entity"
)

type GetSalidasPendientes struct {
	Repo salidasDomain.SalidasInterface
}

func (uc *GetSalidasPendientes) Execute(id_cendis int32) (*[]salidaEntity.SalidaEntity, error) {
	return uc.Repo.GetSalidasPendientes(id_cendis)
}
