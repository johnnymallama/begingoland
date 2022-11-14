package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/johnnymallama/begingoland/bd"
)

func LeoTweetSeguidor(w http.ResponseWriter, r *http.Request) {
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
	data, _, err := bd.LeoTweetSeguidor(IDUsuario, int64(paginaInt))
	if err != nil {
		http.Error(w, "Error al consultar listado de usuarios "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	w.WriteHeader(http.StatusCreated)
}
