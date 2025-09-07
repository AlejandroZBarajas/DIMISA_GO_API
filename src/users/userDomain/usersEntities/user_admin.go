package usersEntities

type UserAdminEntity struct{
	UserEntity
	Id_user int32 `json:"id_user"`
}