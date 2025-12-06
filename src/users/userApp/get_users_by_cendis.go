package userApp

import (
	"DIMISA/src/users/userDomain"
	"DIMISA/src/users/userDomain/usersEntities"
)

type GetUsersByCendisUseCase struct {
	Repo userDomain.UserInterface
}

func (uc *GetUsersByCendisUseCase) Execute(cendisID int32) ([]*usersEntities.UserCendisEntity, error) {
	return uc.Repo.GetByCendisID(cendisID)
}
