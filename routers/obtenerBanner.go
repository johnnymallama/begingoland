package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/johnnymallama/begingoland/bd"
)

func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	usuarioId := r.URL.Query().Get("id")
	if len(usuarioId) < 1 {
		http.Error(w, "Debe enviar el id de usuario ", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscoPerfil(usuarioId)
	if err != nil {
		http.Error(w, "Usuario no encontrado ", http.StatusBadRequest)
		return
	}
	openFile, err := os.Open("uploads/banners/" + perfil.Banner)
	if err != nil {
		http.Error(w, "Banner no encontrado ", http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "Error al copiar archivo ", http.StatusBadRequest)
		return
	}
}
