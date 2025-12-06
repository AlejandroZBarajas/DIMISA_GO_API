package entradasApp

import (
	"DIMISA/src/entradas/entradasDomain"
	"DIMISA/src/entradas/entradasDomain/entradaEntity"
)

type CapturarEntradaUseCase struct {
	Repo entradasDomain.EntradaInterface
}

func (uc *CapturarEntradaUseCase) Execute(entrada *entradaEntity.EntradaRequest) error {
	return uc.Repo.CapturarEntrada(entrada)
}
