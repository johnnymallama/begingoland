package routers

import (
	"errors"
	"strings"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/johnnymallama/begingoland/bd"
	"github.com/johnnymallama/begingoland/jwt"
	"github.com/johnnymallama/begingoland/models"
)

var Email, IDUsuario string

func ProcesoToken(token string) (*models.Claim, bool, string, error) {
	miClave := jwt.KeyToken()
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, "", errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwtgo.ParseWithClaims(token, claims, func(t *jwtgo.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("token invalido")
	}
	return claims, false, "", err
}
