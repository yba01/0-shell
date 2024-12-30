package file_commands

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Fonction pour gerer les commande

func HandleCommand(intput string) {
	args := strings.Fields(intput) // Découpe l'entrée en mots (séparés par des espaces)

	switch args[0] {
	case "cat":
		cat(args[1:])
	case "cp":
		if len(args) != 3 {
			fmt.Println("Usage: cp source destination")
			return
		}
		if err := cp(args[1], args[2]); err != nil {
			fmt.Println(err)
			return
		}
	case "mv":
		if len(args) < 3 {
			fmt.Println("Usage: mv source destination")
			return
		}
		if err := mv(args[1], args[2]); err != nil {
			fmt.Println(err)
			return
		}
	default:
	}

}

func cat(files []string) {
	if len(files) == 0 {
		fmt.Println("Error in arguments; ex: cat your_file")
		return
	}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("Error opening file %s: %v\n", file, err)
			continue
		}
		defer f.Close()

		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
		}
	}
}

func cp(src, dest string) error {
	info, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("source error: %v", err)
	}

	if info.IsDir() {
		return copyDir(src, dest)
	}
	return copyFile(src, dest)
}

func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func copyDir(srcDir, destDir string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, relPath)
		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}
		return copyFile(path, destPath)
	})
}

func mv(src, dest string) error {
	err := os.Rename(src, dest)
	if err != nil {
		return err
	}
	return nil
}