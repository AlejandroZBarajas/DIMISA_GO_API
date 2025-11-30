package salidasApp

import (
	"DIMISA/src/salidas/salidasDomain"
)

type DeleteSalida struct {
	Repo salidasDomain.SalidasInterface
}

func (uc *DeleteSalida) Execute(id_salida int32) error {
	return uc.Repo.DeleteSalida(id_salida)
}
