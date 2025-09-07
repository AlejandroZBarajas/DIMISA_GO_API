package userDomain

import "DIMISA/src/users/userDomain/usersEntities"

type UsersEnfInterface interface {
	AsignarArea(userid int32, areaid int32) error

	GetByAreaID(areaid int32) ([]*usersEntities.UserEnfermeriaEntity, error)
}
