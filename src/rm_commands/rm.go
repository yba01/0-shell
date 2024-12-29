package rm_commands

import (
	"fmt"
	"os"
	"strings"

)

func HandleCommand(intput string) {
	args := strings.Fields(intput) // Découpe l'entrée en mots (séparés par des espaces)

	switch args[0] {
	case "rm":
		err := RemoveFiles(args[1:]) // Appel de RemoveFiles avec les arguments
		if err != nil {
			fmt.Println(err) // Afficher l'erreur si RemoveFiles échoue
		}
	default:
	}

}

func FormatError(msg string, args ...interface{}) error {
	return fmt.Errorf("\x1b[31m"+msg+"\x1b[0m", args...)
}

func RemoveFiles(args []string) error {
	if len(args) == 0 {
		return FormatError("rm: missing operand")
	}

	recursive := false
	files := []string{}

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			if arg == "-r" {
				recursive = true
			} else {
				// Afficher une erreur pour les options invalides et retourner immédiatement
				return FormatError("rm: invalid option -- %s", arg)
			}
		} else {
			files = append(files, arg) // Ajouter uniquement les fichiers valides
		}
	}

	// Vérifier si des fichiers ont été ajoutés
	if len(files) == 0 {
		return FormatError("rm: missing file operand")
	}

	for _, file := range files {
		err := remove(file, recursive)
		if err != nil {
			return FormatError("rm: %v", err)
		}
	}

	return nil
}

func remove(path string, recursive bool) error {
	info, err := os.Stat(path)
	if err != nil {
		return FormatError("rm: %v", err)	}

	if !info.IsDir() {
		err = os.Remove(path)
		if err != nil {
			return FormatError("rm: %v", err)		}
		return nil
	}

	if !recursive {
		return FormatError("rm: cannot remove '%s': Is a directory", path)	}

	err = os.RemoveAll(path)
	if err != nil {
		return FormatError("rm: %v", err)	}
	return nil
}

func notAllowed(str string) (bool, string) {
	if len(str) > 1 {
		for _, char := range str[1:] {
			if !strings.Contains("r", string(char)) {
				return false, string(char)
			}
		}
		return true, ""
	}
	return false, str
}