package userApp

import (
	"DIMISA/src/users/userDomain"
	"DIMISA/src/users/userDomain/usersEntities"
)

type CreateUserUseCase struct {
	Repo userDomain.UserInterface
}

func (uc *CreateUserUseCase) Execute(user *usersEntities.UserEntity) (int32, error) {
	return uc.Repo.CreateUser(user)
}
