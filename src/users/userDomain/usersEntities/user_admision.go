package usersEntities

type UserAdmisionEntity struct{
	UserEntity
	Id_user int32 `json:"id_user"`
}