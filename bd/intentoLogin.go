package bd

import (
	"github.com/johnnymallama/begingoland/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, pass string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)

	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(pass)
	passwordBytesBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBytesBD, passwordBytes)

	if err != nil {
		return usu, false
	}
	return usu, true
}
