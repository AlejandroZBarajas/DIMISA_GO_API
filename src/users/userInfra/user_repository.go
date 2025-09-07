package userInfra


import (
	"golang.org/x/crypto/bcrypt"
    "fmt"
	"database/sql"
	"DIMISA/src/users/userDomain/usersEntities"
	//"DIMISA/src/users/userDomain"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user *usersEntities.UserEntity) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return fmt.Errorf("error al hashear la contraseña: %v", err)
    }
	query := `INSERT INTO usuarios (nombres, apellido1, apellido2, username, password, id_rol) 
	          VALUES (?, ?, ?, ?, ?, ?)`
	_, err = r.DB.Exec(query, user.Nombres, user.Apellido1, user.Apellido2, user.Username, string(hashedPassword), user.Id_rol)
	return err
}