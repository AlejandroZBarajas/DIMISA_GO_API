package auth

import (
	"DIMISA/src/users/userDomain/usersEntities"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
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
	IdUsuario     int32  `json:"id_usuario"`
	IdRol         int32  `json:"id_rol"`
	Token         string `json:"token"`
	NombreUsuario string `json:"nombre_usuario"`
	IdArea        *int32 `json:"id_area,omitempty"`   // Solo para rol 5
	IdCendis      *int32 `json:"id_cendis,omitempty"` // Solo para rol 6
}

func (lh *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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

	// Comparar password
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)) != nil {
		http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
		return
	}

	// Leer el secreto JWT
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		http.Error(w, "JWT_SECRET no definido", http.StatusInternalServerError)
		return
	}

	// Claims del token
	claims := jwt.MapClaims{
		"id_usuario": user.Id_usuario,
		"id_rol":     user.Id_rol,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		http.Error(w, "Error generando token", http.StatusInternalServerError)
		return
	}

	// ------------------------------------------------------------------
	// Obtener datos extra según el rol
	// ------------------------------------------------------------------
	var idArea *int32
	var idCendis *int32

	switch user.Id_rol {
	case 5: // Enfermería
		queryArea := `SELECT id_area FROM enfermeria_users WHERE id_user = ?`
		_ = lh.DB.QueryRow(queryArea, user.Id_usuario).Scan(&idArea)

	case 6: // Cendis
		queryCendis := `SELECT id_cendis FROM unidosis_users WHERE id_user = ?`
		_ = lh.DB.QueryRow(queryCendis, user.Id_usuario).Scan(&idCendis)
	}

	nombreCompleto := strings.TrimSpace(
		fmt.Sprintf("%s %s %s", user.Nombres, user.Apellido1, user.Apellido2),
	)
	// ------------------------------------------------------------------
	// Respuesta final
	// ------------------------------------------------------------------
	resp := loginResponse{
		IdUsuario:     user.Id_usuario,
		IdRol:         user.Id_rol,
		Token:         tokenString,
		NombreUsuario: nombreCompleto,
		IdArea:        idArea,
		IdCendis:      idCendis,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
