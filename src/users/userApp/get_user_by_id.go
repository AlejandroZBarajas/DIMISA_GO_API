package userApp

import (
	"DIMISA/src/users/userDomain"
	"DIMISA/src/users/userDomain/usersEntities"
)

type GetUserByIDUseCase struct {
	Repo userDomain.UserInterface
}

func (uc *GetUserByIDUseCase) Execute(id int32) (*usersEntities.UserEntity, error) {
	return uc.Repo.GetById(id)
}
