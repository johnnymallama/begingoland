package routers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/johnnymallama/begingoland/bd"
	"github.com/johnnymallama/begingoland/models"
)

func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, _ := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}
	var usuario models.Usuario
	usuario.Avatar = IDUsuario + "." + extension
	status, err := bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil {
		http.Error(w, "Error al grabar el avatar "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
	w.WriteHeader(http.StatusCreated)
}
