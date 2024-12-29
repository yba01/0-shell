package rm_commands

import (
	"fmt"
	"os"
	"strings"

)

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
			if strings.Contains(arg, "r") {
				recursive = true
			}
		} else {
			files = append(files, arg)
		}
	}

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