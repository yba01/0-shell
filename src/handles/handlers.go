package tools

import (
	file_cmd "0-shell/src/file_commands"
	int_cmd "0-shell/src/internal_commands"
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func IsValidCommand(command string) bool {
	commands := []string{"echo", "cd", "ls", "pwd", "cat", "cp", "rm", "mv", "mkdir", "exit"}
	// Vérifie si la commande est dans la liste des commandes autorisées
	isValidCommand := false
	for _, cmd := range commands {
		if strings.HasPrefix(command, cmd) { // Vérifie si la commande commence par un élément de la liste
			isValidCommand = true
			break
		}
	}
	return isValidCommand
}

func Loop() {
	// repertoire actuel
	var current_dir string
	// Canal pour capturer les signaux SIGINT
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT)

	// Goroutine pour gérer SIGINT
	go func() {
		for range sigChannel {
			// Message affiché lors de l'appui sur Ctrl + C
			fmt.Printf("\n%s$ ", current_dir)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {

		dir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			return
		}
		// Mis a jour du repertoire pour l'appuit sur Ctrl + C
		current_dir = dir

		fmt.Printf("%s$ ", dir)
		input, err := reader.ReadString('\n') // Lit la commande jusqu'a appuyer sur entrée
		if err != nil {
			if err == io.EOF { // Quand on appuit sur Ctrl + D
				fmt.Println("\nexit")
				return
			}
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.ReplaceAll(input, "\"", "")

		command := strings.TrimSpace(input) // Supprime les espaces inutiles
		// verifier si la commande est dans la liste des commandes

		// Verifie si la commande est vide
		if command == "" {
			continue
		}

		if !IsValidCommand(command) {
			fmt.Printf("Command '%s' not found\n", command)
			continue
		}

		int_cmd.HandleCommand(command)
		file_cmd.HandleCommand(command)
	}
}
