package userApp

import (
	"DIMISA/src/users/userDomain"
	"DIMISA/src/users/userDomain/usersEntities"
)

type GetAllUsersUseCase struct {
	Repo userDomain.UserInterface
}

func (uc *GetAllUsersUseCase) Execute() ([]*usersEntities.UserEntity, error) {
	return uc.Repo.GetAll()
}
