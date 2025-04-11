# Game-projet-bot

---

### Description du script

Ce script en Go permet de charger un **token Discord** depuis une variable d'environnement nommée `DISCORD_TOKEN`.

- Il utilise la fonction `os.Getenv("DISCORD_TOKEN")` pour récupérer la valeur du token.
- Si le token est manquant (la variable d'environnement n'est pas définie), il affiche le message `Token manquant.`.
- Si le token est trouvé, il peut être utilisé pour se connecter à l'API Discord, mais ce code ne l'utilise pas encore dans cet exemple.

### Comment utiliser le script

1. **Définir la variable d'environnement `DISCORD_TOKEN`** :
   - Sur Windows : Suivez les étapes ci-dessus pour ajouter une nouvelle variable d'environnement.
   - Sur macOS/Linux : Ajoutez la ligne suivante dans votre fichier `~/.bashrc` ou `~/.zshrc` :

     ```bash
     export DISCORD_TOKEN="votre_token_ici"
     ```
     Puis, exécutez `source ~/.bashrc` ou `source ~/.zshrc` pour charger la nouvelle variable.

2. **Lancer le script** :
   - Assurez-vous que Go est installé sur votre machine.
   - Ouvrez un terminal et exécutez le script Go avec la commande :

     ```bash
     go run votre_script.go
     ```

Si le token est bien défini, le script pourra l'utiliser pour effectuer des opérations (par exemple, se connecter à un bot Discord).

---