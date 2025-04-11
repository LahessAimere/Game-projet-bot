package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// DiscordBot struct contenant la session du bot
type DiscordBot struct {
	session *discordgo.Session
}

// NewDiscordBot crée une nouvelle instance de DiscordBot avec le token
func NewDiscordBot(token string) (*DiscordBot, error) {
	// Crée une nouvelle session Discord avec le token fourni
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la création de la session: %v", err)
	}

	// Définir les intentions pour accéder aux messages du serveur
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	// Retourne l'objet DiscordBot avec la session initialisée
	return &DiscordBot{session: dg}, nil
}

// Run démarre le bot et attend une interruption pour l'arrêter
func (bot *DiscordBot) Run() error {
	// Ouverture de la connexion à Discord
	if err := bot.session.Open(); err != nil {
		return fmt.Errorf("erreur lors de l'ouverture de la connexion: %v", err)
	}

	// Message indiquant que le bot est en ligne
	fmt.Println("Bot en ligne ! Tapez Ctrl+C pour quitter.")

	// Attendre l'interruption (Ctrl+C ou signal système)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	// Message indiquant que le bot est hors ligne
	fmt.Println("Le bot est maintenant hors ligne.")

	// Fermer la session Discord proprement
	if err := bot.session.Close(); err != nil {
		return fmt.Errorf("erreur lors de la fermeture de la session: %v", err)
	}

	return nil
}
