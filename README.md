# 0-shell

Le projet 0-shell consiste à créer une mini interface shell Unix capable d'exécuter des commandes courantes. Le but est d'apprendre les fondamentaux d'un shell tout en explorant les concepts clés des systèmes d'exploitation Unix, tels que la création et la gestion des processus.

# Objectifs

1. Créer un shell minimaliste qui :

   - Affiche un prompt ($) et attend des commandes.
- Exécute les commandes saisies de manière séquentielle.
- Gère les erreurs de manière claire et informative.
- Implémente un ensemble de commandes Unix de base, comme echo, ls, cd, etc.
- Ne dépend pas des commandes système externes (elles doivent être codées depuis zéro).

2. Approfondir les notions suivantes :

- Interaction avec le système d'exploitation (API Unix).
- Création et synchronisation des processus.
- Manipulation des entrées/sorties.
- Respect des bonnes pratiques de développement logiciel.
- Favoriser le travail d’équipe, avec une répartition des tâches cohérente et efficace.

# Fonctionnalités demandées
 Votre shell doit prendre en charge les commandes suivantes, toutes codées en interne :

1. echo : Afficher du texte.
2. cd : Changer le répertoire courant.
3. ls : Lister les fichiers avec options -l, -a, -F.
4. pwd : Afficher le chemin du répertoire courant.
5. cat : Lire et afficher le contenu d’un fichier.
6. cp : Copier un fichier ou un répertoire.
7. rm : Supprimer un fichier ou répertoire avec l’option -r.
8. mv : Déplacer ou renommer un fichier.
9. mkdir : Créer un nouveau répertoire.
10. exit : Quitter le shell.

Le shell doit aussi :

- Gérer les interruptions avec Ctrl + D.
- Fournir des messages d'erreur clairs.
