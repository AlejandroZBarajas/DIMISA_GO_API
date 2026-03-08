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
	IdArea        *int32 `json:"id_area,omitempty"` // Solo para rol 5
	Area          string `json:"area"`
	IdCendis      *int32 `json:"id_cendis,omitempty"` // Solo para rol 6
	Cendis        string `json:"cendis"`
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

	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)) != nil {
		http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
		return
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		http.Error(w, "JWT_SECRET no definido", http.StatusInternalServerError)
		return
	}

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

	var idArea *int32
	var area string
	var idCendis *int32
	var cendis string

	switch user.Id_rol {
	case 5: // Enfermería
		queryIDArea := `SELECT id_area FROM enfermeria_users WHERE id_user = ?`
		if err := lh.DB.QueryRow(queryIDArea, user.Id_usuario).Scan(&idArea); err != nil {
			fmt.Printf("Error obteniendo id_area: %v\n", err)
		}
		if idArea != nil {
			queryArea := `SELECT nombre_area FROM areas WHERE id_area = ?`
			if err := lh.DB.QueryRow(queryArea, idArea).Scan(&area); err != nil {
				fmt.Printf("Error obteniendo nombre_area: %v\n", err)
			}
		}

	case 6: // Cendis
		queryIDCendis := `SELECT id_cendis FROM unidosis_users WHERE id_user = ?`
		if err := lh.DB.QueryRow(queryIDCendis, user.Id_usuario).Scan(&idCendis); err != nil {
			fmt.Printf("Error obteniendo id_cendis: %v\n", err)
		}
		if idCendis != nil {
			queryCendis := `SELECT cendis_nombre FROM cendis WHERE id_cendis = ?`
			if err := lh.DB.QueryRow(queryCendis, idCendis).Scan(&cendis); err != nil {
				fmt.Printf("Error obteniendo cendis_nombre: %v\n", err)
			}
		}
	}

	nombreCompleto := strings.TrimSpace(
		fmt.Sprintf("%s %s %s", user.Nombres, user.Apellido1, user.Apellido2),
	)

	resp := loginResponse{
		IdUsuario:     user.Id_usuario,
		IdRol:         user.Id_rol,
		Token:         tokenString,
		NombreUsuario: nombreCompleto,
		IdArea:        idArea,
		Area:          area,
		IdCendis:      idCendis,
		Cendis:        cendis,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
