package salidasApp

import (
	"DIMISA/src/salidas/salidasDomain"
	salidaEntity "DIMISA/src/salidas/salidasDomain/entity"
)

type CreateSalida struct {
	Repo salidasDomain.SalidasInterface
}

func (uc *CreateSalida) Execute(salida *salidaEntity.SalidaEntity) error {
	return uc.Repo.CreateSalida(salida)
}
