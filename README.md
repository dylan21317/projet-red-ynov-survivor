🎓 Ynov RPG - Le Campus Tactique
Un jeu de rôle textuel (CLI) immersif développé en Go (Golang), se déroulant au cœur d'un campus high-tech impitoyable. Incarnez un survivant, explorez le terrain, gérez votre inventaire et affrontez les mentors pour décrocher votre diplôme !

📖 À propos du Projet
Ynov Lost colony est un projet complet développé en Go, mettant en avant la modularité du code à travers plusieurs paquets ou fichiers source. Le joueur y fait face à des ennemis robotiques et mutants dans un environnement universitaire cyberpunk. Le jeu intègre des mécaniques poussées de gestion de classe, des combats graphiques en ASCII, un système d'artisanat via une imprimante 3D et une exploration en temps réel sur une grille interactive .
🚀 Fonctionnalités Clés
3 Classes de Personnages Distinctes :
👤 Humain : Profil équilibré, idéal pour débuter (, , ).
🧝 Elfe : Profil agile et orienté magie puissante (, , ).
🛡️ Nain : Gros tank résistant au corps à corps (, , ).
Système de Combat Tactique (Tour par Tour) :
Animations dynamiques en art ASCII dans le terminal.
Gestion des temps de recharge (cooldowns) des sorts et du coût en énergie magique ().
Altérations d'état dynamiques (brûlures sur la durée, étourdissements).
Exploration Libre () :
Déplacement en temps réel avec les touches Z, S, Q, D.
Événements aléatoires sur la carte : monstres agressifs, casiers de butin et pièges électriques.
Mécanique de cooldown d'exploration pour rythmer l'aventure.
Économie, Marchand & Imprimante 3D :
Boutique pour acheter des consommables et de nouveaux sorts actifs (Boule de Feu, Jet de Café Brûlant, Ctrl+Alt+Suppr).
Récupération de matériaux sur les monstres (Plume de Corbeau, Cuir de Sanglier, Peau de Troll) pour fabriquer de l'équipement légendaire sur l'imprimante 3D.
Boss Final (Les 3 Mentors) :
Épreuve suprême débloquée après avoir prouvé votre valeur (requis : au moins 3 victoires).

📂 Architecture du Code

Le projet est divisé en plusieurs fichiers Go pour assurer une séparation claire des responsabilités :
Fichier
Rôle principal
main1.go
Point d'entrée du jeu, gestion des menus principaux, du marchand, de l'imprimante 3D et des boucles de fin de jeu.
combat.go
Gestion de la boucle de combat tactique, des attaques, de l'affichage de l'arène ASCII et des altérations d'état.
player.go
Gestion des statistiques du personnage, de la progression d'expérience, des niveaux, des équipements et de l'inventaire.
world.go
Logique d'exploration de la grille interactive  et gestion des événements de terrain.
types.go
Définition des structures globales (Character, Enemy, Skill, Equipment) et des codes couleurs ANSI.
ascii.txt
Fichiers de ressources graphiques pour les sprites du boss et du mentor suprême.

🛠️ Prérequis & Dépendances
Langage Go : Version 1.18 ou supérieure recommandée.
Terminal compatible ANSI : Un terminal moderne (Linux/macOS Terminal, Windows Terminal, VS Code Integrated Terminal) pour un affichage correct des couleurs et des sprites ASCII.
Dépendance externe : Le paquet de gestion du terminal pour les déplacements en temps réel.
📦 Installation & Lancement
Cloner ou placer l'ensemble des fichiers (main1.go, combat.go, player.go, world.go, types.go, ascii.txt) dans un même répertoire de travail.
Installer la dépendance requise pour la gestion des touches en mode brut :
go get golang.org/x/term


Compiler et lancer le jeu directement via la ligne de commande Go :
go run .


🎮 Guide de Jeu
Navigation dans les menus : Tapez le chiffre correspondant à l'action désirée et validez avec la touche Entrée.
Exploration du Campus : Utilisez les touches de direction clavier :
Z : Aller vers le Haut
S : Aller vers le Bas
Q : Aller vers la Gauche
D : Aller vers la Droite
En combat : Utilisez vos compétences stratégiques, surveillez vos points de vie () et de magie (), ou utilisez vos potions en cas de coup dur.
