package internal_commands

import (
	"fmt"
	"os"
	"strings"
)

// Fonction pour gerer les commande

func HandleCommand(intput string) {
	args := strings.Fields(intput) // Découpe l'entrée en mots (séparés par des espaces)

	switch args[0] {
	case "exit":
		exitShell()
	case "echo":
		echoCommand(args)
	case "pwd":
		pwdCommand()
	default:
	}

}

// Implémentation de la commande `exit`
func exitShell() {
	fmt.Println("Exiting shell...")
	os.Exit(0)
}

// Implémentation de la commande `echo`
func echoCommand(args []string) {
	// Joindre les arguments et les afficher
	fmt.Println(strings.Join(args[1:], " "))
}

// Implémentation de la commande `pwd`
func pwdCommand() {
	cwd, err := os.Getwd() // Récupérer le répertoire courant
	if err != nil {
		fmt.Println("Error retrieving current directory")
	} else {
		fmt.Println(cwd) // Afficher le répertoire courant
	}
}
