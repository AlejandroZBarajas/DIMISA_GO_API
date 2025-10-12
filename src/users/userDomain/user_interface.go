package userDomain

import "DIMISA/src/users/userDomain/usersEntities"

type UserInterface interface {
	CreateUser(user *usersEntities.UserEntity) (id int32, err error)

	CreateUserCendis(idUser, idCendis int32) error

	CreateUserEnfermeria(idUser, idArea int32) error

	CreateAdminUser(userID int32) error
	CreateJefeUser(userID int32) error
	CreateAdmisionUser(userID int32) error

	UpdateUser(user *usersEntities.UserEntity) error

	GetById(id int32) (*usersEntities.UserEntity, error)

	GetByRol(rol int32) ([]*usersEntities.UserEntity, error)

	GetAll() ([]*usersEntities.UserEntity, error)

	DeleteUser(id int32) error

	//ExistUsername(username string) (bool, error)

	GetByAreaID(id int32) ([]*usersEntities.UserEnfermeriaEntity, error)

	GetByCendisID(id int32) ([]*usersEntities.UserCendisEntity, error)
}
