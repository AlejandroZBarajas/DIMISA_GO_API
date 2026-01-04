package salidasApp

import (
	"DIMISA/src/salidas/salidasDomain"
	salidaEntity "DIMISA/src/salidas/salidasDomain/entity"
)

type AddToSalida struct {
	Repo salidasDomain.SalidasInterface
}

func (uc *AddToSalida) Execute(id_cendis, id_area, tipo int32, claves *[]salidaEntity.SalidaDetalleEntity) error {
	return uc.Repo.AddToSalida(id_cendis, id_area, tipo, claves)
}
