package main

import (
	"github.com/MatiasCermak/go-meli-integration/pkg/router"
)

type Prueba struct {
	Nombre   string
	Apellido string
}

func main() {
	router.Run()
}
