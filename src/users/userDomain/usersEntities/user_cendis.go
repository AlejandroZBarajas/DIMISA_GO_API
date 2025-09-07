package usersEntities

type UserCendisEntity struct {
	UserEntity
	Id_user   int32 `json:"id_user"`
	Id_cendis int32 `json:"id_cendis"`
	//Id_turno int32 `json:"id_turno"`
}

func CreateUserCendis(id_user int32, id_cendis int32 /* id_turno int32 */) *UserCendisEntity {
	return &UserCendisEntity{
		Id_user:   id_user,
		Id_cendis: id_cendis,
		//Id_turno: id_turno,
	}
}

/*
EJEMPLO de uso

cendisUser := UserCendisEntity{
	UserEntity: UserEntity{
		Id_usuario: 1,
		Nombres:    "Ana",
		Apellido1:  "López",
		Username:   "ana.l",
		Id_rol:     2,
	},
	Id_cendis: 10,
	Id_turno:  1,
} */
