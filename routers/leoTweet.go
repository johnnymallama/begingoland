package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/johnnymallama/begingoland/bd"
)

func LeoTweet(w http.ResponseWriter, r *http.Request) {
	pagina := r.URL.Query().Get("pagina")
	if len(pagina) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	paginaInt, err := strconv.Atoi(pagina)
	if err != nil {
		http.Error(w, "Error en valor de pagina debe ser numerico "+err.Error(), http.StatusBadRequest)
		return
	}

	pag := int64(paginaInt)
	data, status, err := bd.LeoTweet(IDUsuario, pag)
	if err != nil {
		http.Error(w, "Error consultando tweets "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado consultar los tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
