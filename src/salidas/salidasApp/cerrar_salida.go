package salidasApp

import (
	"DIMISA/src/salidas/salidasDomain"
)

type CerrarSalida struct {
	Repo salidasDomain.SalidasInterface
}

func (uc *CerrarSalida) Execute(id_salida int32) error {
	return uc.Repo.CerrarSalida(id_salida)
}
