package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/johnnymallama/begingoland/bd"
	"github.com/johnnymallama/begingoland/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), http.StatusBadRequest)
		return
	}
	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al registrar el Tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se logro guardar el Tweet", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
