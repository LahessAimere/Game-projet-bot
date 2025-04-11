package main

import (
	"fmt"
	"os"
)

func main() {
	// Charge le token depuis les variables d'environnement
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		fmt.Println("Token manquant.")
		return
	}
}
