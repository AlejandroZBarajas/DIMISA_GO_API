package userApp

import (
	"DIMISA/src/users/userDomain"
	"DIMISA/src/users/userDomain/usersEntities"
)

type UpdateUserUseCase struct {
	Repo userDomain.UserInterface
}

func (uc *UpdateUserUseCase) Execute(user *usersEntities.UserEntity) error {
	return uc.Repo.UpdateUser(user)
}
