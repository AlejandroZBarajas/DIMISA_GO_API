package userDomain

import "DIMISA/src/users/userDomain/usersEntities"

type UsersCendisInterface interface {

	AsignarCendis(userid int32, cendisid int32) error

	AsignarTurno(userid int32, turnoid int32) error

	GetByCendisID(cendisid int32)([]*usersEntities.UserCendisEntity, error)

}

