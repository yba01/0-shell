package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	int_cmd "0-shell/src/internal_commands"
)

func main() {
	fmt.Println("Hello, Go!")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n') //Lit la commande jusqu'a appuyer sur entrée
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		command := strings.TrimSpace(input) // Supprime les espaces inutiles

		// Verifie si la commande est vide
		if command == ""{
			continue
		}

		int_cmd.HandleCommand(command)
	}
}
