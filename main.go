package main

import (
	"fmt"
	"log"

	"github.com/BrandokVargas/dukibot/bot"
)

func main() {
	fmt.Println("Iniciando bot...")

	dg, err := bot.Start()
	if err != nil {
		log.Fatalf("Error iniciando el bot: %v", err)
	}

	// Mantener el bot corriendo
	defer dg.Close()
	select {} // Bloquea el programa para que no termine
}
