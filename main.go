package main

import (
	"log"

	"github.com/johnnymallama/begingoland/bd"
	"github.com/johnnymallama/begingoland/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
