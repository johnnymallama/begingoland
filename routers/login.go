package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/johnnymallama/begingoland/bd"
	"github.com/johnnymallama/begingoland/jwt"
	"github.com/johnnymallama/begingoland/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos"+err.Error(), http.StatusBadRequest)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", http.StatusBadRequest)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "Usuario y/o contraseña invalidos", http.StatusBadRequest)
		return
	}
	jwtkey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Occurio un error al intentar generar el token correspondiente "+err.Error(), http.StatusBadRequest)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtkey,
	}
	w.Header().Set("Content-Type", "application/json")
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtkey,
		Expires: expirationTime,
	})
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
