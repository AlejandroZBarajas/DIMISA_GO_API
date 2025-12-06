package userApp

import (
	"DIMISA/src/users/userDomain"
	"DIMISA/src/users/userDomain/usersEntities"
)

type GetUsersByAreaUseCase struct {
	Repo userDomain.UserInterface
}

func (uc *GetUsersByAreaUseCase) Execute(areaID int32) ([]*usersEntities.UserEnfermeriaEntity, error) {
	return uc.Repo.GetByAreaID(areaID)
}
