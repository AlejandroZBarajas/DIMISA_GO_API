package auth

import (
	"DIMISA/src/users/userDomain/usersEntities"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	DB *sql.DB
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	IdUsuario int32  `json:"id_usuario"`
	IdRol     int32  `json:"id_rol"`
	Token     string `json:"token"`
}

func (lh *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		//	fmt.Printf("Password recibida: '%s'\n", req.Password)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	var user usersEntities.UserEntity
	var hashedPassword string

	query := `SELECT id_usuario, nombres, apellido1, apellido2, username, password, id_rol
              FROM usuarios WHERE username = ?`
	err := lh.DB.QueryRow(query, req.Username).Scan(
		&user.Id_usuario,
		&user.Nombres,
		&user.Apellido1,
		&user.Apellido2,
		&user.Username,
		&hashedPassword,
		&user.Id_rol,
	)
	if err == sql.ErrNoRows {
		http.Error(w, "Usuario no encontrado", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Error DB: %v", err), http.StatusInternalServerError)
		return
	}

	/* fmt.Printf("Hash desde DB: '%s'\n", hashedPassword)
	fmt.Printf("Password recibida: '%s'\n", req.Password) */
	errCmp := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	fmt.Printf("Resultado bcrypt: %v\n", errCmp)

	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)) != nil {
		http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
		return
	}

	// Leer el secreto desde .env
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		http.Error(w, "JWT_SECRET no definido", http.StatusInternalServerError)
		return
	}

	// Crear token JWT
	claims := jwt.MapClaims{
		"id_usuario": user.Id_usuario,
		"id_rol":     user.Id_rol,
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		http.Error(w, "Error generando token", http.StatusInternalServerError)
		return
	}

	resp := loginResponse{
		IdUsuario: user.Id_usuario,
		IdRol:     user.Id_rol,
		Token:     tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
