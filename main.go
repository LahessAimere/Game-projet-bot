package main

import (
	"fmt"
	"os"

	"Game-projet-bot/bot"
)

func main() {
	// Charge le token depuis les variables d'environnement
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		fmt.Println("Token manquant. Veuillez définir la variable d'environnement DISCORD_TOKEN.")
		return
	}

	// Crée et démarre le bot
	b, err := bot.NewDiscordBot(token)
	if err != nil {
		fmt.Println("Erreur lors de la création du bot:", err)
		return
	}

	// Exécution du bot
	if err := b.Run(); err != nil {
		fmt.Println("Erreur lors de l'exécution du bot:", err)
	}
}
