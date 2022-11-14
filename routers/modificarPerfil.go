package routers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/johnnymallama/begingoland/bd"
	"github.com/johnnymallama/begingoland/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error en lectura de datos "+err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(data, &t)
	if err != nil {
		http.Error(w, "Datos Inconrrectos "+err.Error(), http.StatusBadRequest)
		return
	}

	status, err := bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro, reintente nuevamente "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario", http.StatusBadRequest)
		return
	}
}
