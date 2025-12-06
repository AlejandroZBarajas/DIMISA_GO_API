package userApp

import (
	"DIMISA/src/users/userDomain"
	"DIMISA/src/users/userDomain/usersEntities"
)

type GetUsersByRolUseCase struct {
	Repo userDomain.UserInterface
}

func (uc *GetUsersByRolUseCase) Execute(rol int32) ([]*usersEntities.UserEntity, error) {
	return uc.Repo.GetByRol(rol)
}
