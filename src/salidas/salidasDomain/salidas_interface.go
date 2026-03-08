package salidasDomain

import (
	salidaEntity "DIMISA/src/salidas/salidasDomain/entity"
)

type SalidasInterface interface {
	CreateSalida(salida *salidaEntity.SalidaEntity) (int32, error)
	UpdateSalida(id_salida int32, salida *salidaEntity.SalidaEntity) error
	DeleteSalida(id_salida int32) error
	GetSalidasByCendis(id_cendis int32) (*[]salidaEntity.SalidaEntity, error)
	GetSalidasPendientes(id_cendis int32) (*[]salidaEntity.SalidaEntity, error)
	AddToSalida(id_cendis, id_area, tipo int32, claves *[]salidaEntity.SalidaDetalleEntity) error
	CerrarSalida(id_salida int32) error
}
