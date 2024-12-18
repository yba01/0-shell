package main

import (
	file_cmd "0-shell/src/file_commands"
	"bufio"
	"fmt"
	"os"
	"strings"
	// int_cmd "0-shell/src/internal_commands"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			return
		}
		fmt.Printf("%s$ ", dir)
		input, err := reader.ReadString('\n') //Lit la commande jusqu'a appuyer sur entrée
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.ReplaceAll(input, "\"", "")

		command := strings.TrimSpace(input) // Supprime les espaces inutiles

		// Verifie si la commande est vide
		if command == "" {
			continue
		}

		// int_cmd.HandleCommand(command)
		file_cmd.HandleCommand(command)
	}
}
