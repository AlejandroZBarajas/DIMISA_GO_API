package salidasApp

import (
	"DIMISA/src/salidas/salidasDomain"
	salidaEntity "DIMISA/src/salidas/salidasDomain/entity"
)

type UpdateSalida struct {
	Repo salidasDomain.SalidasInterface
}

func (uc *UpdateSalida) Execute(salida *salidaEntity.SalidaEntity) error {
	return uc.Repo.UpdateSalida(salida)
}
