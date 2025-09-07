package usersEntities

type UserEntity struct {

	Id_usuario int32 `json:"id_usuario"`
	Nombres string `json:"nombres"`
	Apellido1       string  `json:"apellido1"`
	Apellido2     string `json:"apellido2"`
	Username string `json:"username"`
	Password   string `json:"-" db:"password"` 
	Id_rol int32 `json:"id_rol"`

}

func CreateUser(
	id_usuario int32,
	nombres string,
	apellido1 string,
	apellido2 string,
	username string,
	password string,
	id_rol int32,
) *UserEntity {
	return &UserEntity{
		Id_usuario: id_usuario,
		Nombres:    nombres,
		Apellido1:  apellido1,
		Apellido2:  apellido2,
		Username:   username,
		Password:   password,
		Id_rol:     id_rol,
	}
}