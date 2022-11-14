package routers

import (
	"encoding/json"
	"net/http"

	"github.com/johnnymallama/begingoland/bd"
)

func BorroTweet(w http.ResponseWriter, r *http.Request) {
	tweetID := r.URL.Query().Get("id")
	if len(tweetID) < 1 {
		http.Error(w, "Debe enviar el id del tweet", http.StatusBadRequest)
		return
	}

	status, err := bd.BorroTweet(tweetID, IDUsuario)
	if err != nil {
		http.Error(w, "Error al eliminar un tweet "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
	w.WriteHeader(http.StatusCreated)
}
