# Game-projet-bot

---

### Description du script

Ce script en Go permet de charger un **token Discord** depuis une variable d'environnement nommée `DISCORD_TOKEN`.

- Il utilise la fonction `os.Getenv("DISCORD_TOKEN")` pour récupérer la valeur du token.
- Si le token est manquant (la variable d'environnement n'est pas définie), il affiche le message `Token manquant. Veuillez définir la variable d'environnement DISCORD_TOKEN.`.
- Si le token est trouvé, il peut être utilisé pour se connecter à l'API Discord.
- Le script crée et exécute un bot Discord en utilisant la bibliothèque `github.com/bwmarrin/discordgo`.
- Le bot restera en ligne jusqu'à ce que l'utilisateur envoie une interruption (Ctrl+C ou signal système).

### Comment utiliser le script

1. **Définir la variable d'environnement `DISCORD_TOKEN`** :
   - **Sur Windows** :
     1. Ouvrez une invite de commandes (cmd).
     2. Tapez la commande suivante pour définir la variable d'environnement :

        ```cmd
        set DISCORD_TOKEN=VotreTokenIci
        ```
     3. Lancez ensuite le programme avec `go run`.

   - **Sur macOS/Linux** :
     1. Ouvrez un terminal.
     2. Ajoutez la ligne suivante dans votre fichier `~/.bashrc` ou `~/.zshrc` :

        ```bash
        export DISCORD_TOKEN="votre_token_ici"
        ```
     3. Rechargez votre fichier de configuration avec la commande suivante :

        ```bash
        source ~/.bashrc   # ou source ~/.zshrc si vous utilisez zsh
        ```
     4. Ensuite, exécutez le script Go avec `go run`.

2. **Lancer le script** :
   - Assurez-vous que Go est installé sur votre machine.
   - Ouvrez un terminal et exécutez le script Go avec la commande :
   
     ```bash
     go run main.go
     ```

Si le token est bien défini, le script pourra l'utiliser pour effectuer des opérations (par exemple, se connecter à un bot Discord).

---

### Structure du projet

- **main.go** : Le point d'entrée du programme, chargé de récupérer le token et de démarrer le bot.
- **bot.go** : Contient la logique pour créer, configurer et exécuter le bot Discord.

---

### Points importants à vérifier

- **Token Discord valide** : Assurez-vous que le token utilisé est valide et que votre bot a été correctement configuré sur le portail développeur de Discord.
- **Permissions du bot** : Vérifiez que votre bot possède les bonnes permissions pour effectuer des actions sur un serveur Discord.
