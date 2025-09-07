package usersEntities

type UserEnfermeriaEntity struct {
    UserEntity
    Id_user  int32 `json:"id_user"`
    Id_area  int32 `json:"id_area"`
    Id_turno int32 `json:"id_turno"`
}

func CreateUserEnfermeria(id_user int32, id_area int32, id_turno int32) *UserEnfermeriaEntity {
	return &UserEnfermeriaEntity{
		Id_user:  id_user,
		Id_area:  id_area,
		Id_turno: id_turno,
	}
}
