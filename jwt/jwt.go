package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/johnnymallama/begingoland/models"
)

func GeneroJWT(u models.Usuario) (string, error) {

	payload := jwt.MapClaims{
		"email":            u.Email,
		"nombre":           u.Nombre,
		"apellidos":        u.Apellidos,
		"fecha_nacimiento": u.FechaNacimiento,
		"biografia":        u.Biografia,
		"ubicacion":        u.Ubicacion,
		"sitioweb":         u.SitioWeb,
		"_id":              u.ID.Hex(),
		"exp":              time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(KeyToken())
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

func KeyToken() []byte {
	return []byte("BeginGolang")
}
