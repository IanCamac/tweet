package main

import (
	"log"

	"github.com/IanCamac/tweet/bd"
	"github.com/IanCamac/tweet/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
