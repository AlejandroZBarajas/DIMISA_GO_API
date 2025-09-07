package userInfra

import (
	"DIMISA/src/users/userDomain/usersEntities"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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

func (r *UserRepository) UpdateUser(user *usersEntities.UserEntity) error {
	// Si el campo Password viene vacío, no actualizamos la contraseña
	if user.Password == "" {
		query := `UPDATE usuarios 
		          SET nombres = ?, apellido1 = ?, apellido2 = ?, username = ?, id_rol = ? 
		          WHERE id_usuario = ?`
		_, err := r.DB.Exec(query, user.Nombres, user.Apellido1, user.Apellido2, user.Username, user.Id_rol, user.Id_usuario)
		return err
	}

	// Si viene la contraseña, la hasheamos y actualizamos todo
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error al hashear la contraseña: %v", err)
	}

	query := `UPDATE usuarios 
	          SET nombres = ?, apellido1 = ?, apellido2 = ?, username = ?, password = ?, id_rol = ? 
	          WHERE id_usuario = ?`
	_, err = r.DB.Exec(query, user.Nombres, user.Apellido1, user.Apellido2, user.Username, string(hashedPassword), user.Id_rol, user.Id_usuario)
	return err
}

func (r *UserRepository) GetAll() ([]*usersEntities.UserEntity, error) {
	query := `SELECT id_usuario, nombres, apellido1, apellido2, username, id_rol 
	          FROM usuarios`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*usersEntities.UserEntity
	for rows.Next() {
		var user usersEntities.UserEntity
		err := rows.Scan(
			&user.Id_usuario,
			&user.Nombres,
			&user.Apellido1,
			&user.Apellido2,
			&user.Username,
			&user.Id_rol,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepository) GetByRol(rol int32) ([]*usersEntities.UserEntity, error) {
	query := `SELECT id_usuario, nombres, apellido1, apellido2, username, id_rol 
	          FROM usuarios WHERE id_rol = ?`

	rows, err := r.DB.Query(query, rol)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*usersEntities.UserEntity
	for rows.Next() {
		var user usersEntities.UserEntity
		err := rows.Scan(
			&user.Id_usuario,
			&user.Nombres,
			&user.Apellido1,
			&user.Apellido2,
			&user.Username,
			&user.Id_rol,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepository) GetById(id int32) (*usersEntities.UserEntity, error) {
	query := `SELECT id_usuario, nombres, apellido1, apellido2, username, id_rol 
	          FROM usuarios WHERE id_usuario = ?`

	var user usersEntities.UserEntity
	err := r.DB.QueryRow(query, id).Scan(
		&user.Id_usuario,
		&user.Nombres,
		&user.Apellido1,
		&user.Apellido2,
		&user.Username,
		&user.Id_rol,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
func (r *UserRepository) DeleteUser(id int32) error {
	query := `DELETE FROM usuarios WHERE id_usuario = ?`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *UserRepository) GetByAreaID(id int32) ([]*usersEntities.UserEnfermeriaEntity, error) {
	query := `SELECT id_usuario, nombres, apellido1, apellido2, username, id_area
	          FROM usuarios_enfermeria WHERE id_area = ?`

	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*usersEntities.UserEnfermeriaEntity
	for rows.Next() {
		var user usersEntities.UserEnfermeriaEntity
		err := rows.Scan(
			&user.Id_usuario,
			&user.Nombres,
			&user.Apellido1,
			&user.Apellido2,
			&user.Username,
			&user.Id_area,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepository) GetByCendisID(id int32) ([]*usersEntities.UserCendisEntity, error) {
	query := `SELECT id_usuario, nombres, apellido1, apellido2, username, id_cendis
	          FROM usuarios_cendis WHERE id_cendis = ?`

	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*usersEntities.UserCendisEntity
	for rows.Next() {
		var user usersEntities.UserCendisEntity
		err := rows.Scan(
			&user.Id_usuario,
			&user.Nombres,
			&user.Apellido1,
			&user.Apellido2,
			&user.Username,
			&user.Id_cendis,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepository) AsignarTurno(userID, turnoID int32) error {
	query := `UPDATE usuarios SET turno_id = ? WHERE id_usuario = ?`
	_, err := r.DB.Exec(query, turnoID, userID)
	return err
}

func (r *UserRepository) AsignarArea(userID, areaID int32) error {
	query := `UPDATE usuarios_enfermeria SET id_area = ? WHERE id_usuario = ?`
	_, err := r.DB.Exec(query, areaID, userID)
	return err
}
