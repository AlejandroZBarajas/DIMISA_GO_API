package userDomain

import "DIMISA/src/users/userDomain/usersEntities"

type UserInterface interface {
	CreateUser(user *usersEntities.UserEntity) error

	UpdateUser(user *usersEntities.UserEntity) error

	GetById(id int32) (*usersEntities.UserEntity, error)

	GetByRol(rol int32) ([]*usersEntities.UserEntity, error)

	GetAll() ([]*usersEntities.UserEntity, error)

	DeleteUser(id int32) error

	//ExistUsername(username string) (bool, error)

	GetByAreaID(id int32) ([]*usersEntities.UserEnfermeriaEntity, error)

	GetByCendisID(id int32) ([]*usersEntities.UserCendisEntity, error)
}
