package salidasDomain

import (
	salidaEntity "DIMISA/src/salidas/salidasDomain/entity"
	//"DIMISA/src/salidas/salidasDomain/salidaEntity"
)

type SalidasInterface interface {
	CreateSalida(salida *salidaEntity.SalidaEntity) error
	UpdateSalida(salida *salidaEntity.SalidaEntity) error
	DeleteSalida(id_salida int32) error
	GetSalidasByCendis(id_cendis int32) (*[]salidaEntity.SalidaEntity, error)
	GetSalidasPendientes(id_cendis int32) (*[]salidaEntity.SalidaEntity, error)
}
