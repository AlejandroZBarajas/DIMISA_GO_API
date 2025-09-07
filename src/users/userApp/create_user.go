package userApp

import (
	"DIMISA/src/users/userDomain/usersEntities"
	"DIMISA/src/users/userDomain"
)

type CreateUserUseCase struct {
	Repo userDomain.UserInterface
}

func (uc *CreateUserUseCase) Execute(user *usersEntities.UserEntity) error {
	return uc.Repo.CreateUser(user)
}
