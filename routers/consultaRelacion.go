package routers

import (
	"encoding/json"
	"net/http"

	"github.com/johnnymallama/begingoland/bd"
	"github.com/johnnymallama/begingoland/models"
)

func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	usuarioRelacionId := r.URL.Query().Get("id")
	if len(usuarioRelacionId) < 1 {
		http.Error(w, "Debe enviar el id del usuario a relacionar", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = usuarioRelacionId

	var respuesta models.RespuestaConsultaRelacion

	status, _ := bd.ConsultoRelacion(t)
	respuesta.Status = status
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
	w.WriteHeader(http.StatusCreated)
}
