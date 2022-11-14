package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/johnnymallama/begingoland/bd"
)

func ListaUsuario(w http.ResponseWriter, r *http.Request) {
	tipo := r.URL.Query().Get("type")
	if len(tipo) < 1 {
		http.Error(w, "Debe enviar el tipo", http.StatusBadRequest)
		return
	}

	pagina := r.URL.Query().Get("page")
	if len(pagina) < 1 {
		http.Error(w, "Debe enviar la paginacion", http.StatusBadRequest)
		return
	}
	paginaInt, err := strconv.Atoi(pagina)
	if err != nil {
		http.Error(w, "Error en parametro paginacion debe ser numerico", http.StatusBadRequest)
		return
	}

	search := r.URL.Query().Get("search")

	data, _, err := bd.LeoUsuarioTodo(IDUsuario, int64(paginaInt), search, tipo)
	if err != nil {
		http.Error(w, "Error al consultar listado de usuarios "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	w.WriteHeader(http.StatusCreated)
}
