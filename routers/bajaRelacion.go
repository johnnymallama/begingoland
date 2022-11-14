package routers

import (
	"encoding/json"
	"net/http"

	"github.com/johnnymallama/begingoland/bd"
	"github.com/johnnymallama/begingoland/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	usuarioRelacionId := r.URL.Query().Get("id")
	if len(usuarioRelacionId) < 1 {
		http.Error(w, "Debe enviar el id del usuario a eliminar relacion", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = usuarioRelacionId

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Error al eliminar relacion de usuario "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
	w.WriteHeader(http.StatusCreated)
}
