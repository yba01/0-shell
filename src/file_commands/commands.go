package file_commands
import (
	"fmt"
	"os"
	"os/user"
	// "path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
)
// Fonction pour gerer les commande
func HandleCommand(intput string) {
	args := strings.Fields(intput) // Découpe l'entrée en mots (séparés par des espaces)
	switch args[0] {
	case "cd":
		changedir(args[1:])
	case "mkdir":
		makedir(args[1:])
	case "ls":
		lsIt(args[1:])
	default:
		// fmt.Printf("Command '%s' not found\n", args[0])
	}
}
// Implémentation de la commande `cd`
func changedir(args []string) {
	var path string
	var err error
	if len(args) > 1 {
		fmt.Println("cd: too many arguments")
		return
	} else if len(args) == 0 {
		// Get the user's home directory
		path, err = os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting home directory: %v\n", err)
			return
		}
	} else if len(args) == 1 {
		path = args[0]
	}
	// Change to another directory
	err = os.Chdir(path)
	if err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		return
	}
}
// Implémentation de la commande `mkdir`
func makedir(args []string) {
	if len(args) == 0 {
		fmt.Println("mkdir: missing operand \nTry 'mkdir --help' for more information.")
	}
	for _, name := range args {
		// Create the directory
		err := os.Mkdir(name, 0755) // 0755 sets permissions for the directory
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
	}
}
// Implémentation de la commande `ls`
func lsIt(args []string) {
	var showAll, longList, classify bool
	var directories []string
	for _, flag := range args {
		if alw, val := notAllowed(flag); !alw {
			if strings.HasPrefix(flag, "-") {
				fmt.Printf("ls: invalid option -- %s\n", val)
				return
			} else {
				directories = append(directories, flag)
			}
		}
		if strings.HasPrefix(flag, "-") && strings.Contains(flag, "a") {
			showAll = true
		}
		if strings.HasPrefix(flag, "-") && strings.Contains(flag, "l") {
			longList = true
		}
		if strings.HasPrefix(flag, "-") && strings.Contains(flag, "F") {
			classify = true
		}
	}
	if len(directories) == 0 {
		directories = append(directories, ".")
	} else {
		sort.Slice(directories, func(i, j int) bool {
			return strings.ToLower(directories[i]) < strings.ToLower(directories[j])
		})
	}
	for i, add := range directories {
		if len(directories) > 1 {
			if add == "." {
				fmt.Printf("%s:\n", add)
			} else {
				fmt.Printf("%s/:\n", add)
			}
		}
		files, totalBlocks, print, err := getFileDir(add, showAll)
		if err != nil {
			fmt.Printf("Error getting file info : %v\n", err)
			return
		}

		// Get current directory
		path, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			return
		}

		if showAll && print {
			dotFile, _ := GetSingleFileInfo(path + "/.")
			dotDotFile, _ := GetSingleFileInfo(path + "/..")
			files = append([]FileInfo{dotFile, dotDotFile}, files...)
			totalBlocks += 16
		}
		// Sort files by name
		sort.Slice(files, func(i, j int) bool {
			return strings.ToLower(files[i].Name) < strings.ToLower(files[j].Name)
		})

		// Filter out hidden files if not showing all
		visibleFiles := []FileInfo{}
		for _, file := range files {
			name := classified(file, classify)
			if showAll || name[0] != '.' {
				visibleFiles = append(visibleFiles, file)
			}
		}

		if longList && print {
			fmt.Printf("total %d\n", totalBlocks/2)
			// Long list format
			for i, file := range visibleFiles {
				name := classified(file, classify)
				mode := file.Mode.String()
				if file.Mode&os.ModeSymlink != 0 {
					mode = "l" + mode[1:]
				}
				links := file.NumLinks
				size := file.Size
				timeStr := formatTime(file.ModTime)

				fmt.Printf("%s %2d %s %s %6d %s %s", mode, links, file.User, file.Group, size, timeStr, name)

				if file.Mode&os.ModeSymlink != 0 {
					target, err := os.Readlink(file.Path)
					if err == nil {
						fmt.Printf(" -> %s", target)
					}
				}
				if i != len(visibleFiles)-1 {
					fmt.Println()
				}
			}
		} else {
			// Column format
			maxTerminalWidth := 80 // Default terminal width
			maxNameLength := 0
			for _, file := range visibleFiles {
				name := classified(file, classify)
				if len(name) > maxNameLength {
					maxNameLength = len(name)
				}
			}

			colWidth := maxNameLength + 2 // Add some padding
			cols := maxTerminalWidth / colWidth
			if cols == 0 {
				cols = 1
			}

			for i := 0; i < len(visibleFiles); i += cols {
				for j := 0; j < cols && (i+j) < len(visibleFiles); j++ {
					name := classified(visibleFiles[i+j], classify)
					fmt.Printf("%-*s", colWidth, name)
				}
				fmt.Println()
			}
		}

		if i != len(directories)-1 {
			fmt.Printf("\n")
		}
	}
}
// Helper function to check if a file is executable
func isExecutable(file FileInfo) bool {
	return file.Mode.Perm()&0111 != 0 // Check execute permissions
}
func notAllowed(str string) (bool, string) {
	if len(str) > 1 {
		for _, char := range str[1:] {
			if !strings.Contains("alF", string(char)) {
				return false, string(char)
			}
		}
		return true, ""
	}
	return false, str
}
func classified(file FileInfo, classify bool) string {
	// Append indicator for -F flag
	if classify {
		if file.IsDir {
			return file.Name + "/"
		} else if isExecutable(file) {
			return file.Name + "*"
		} else {
			return file.Name
		}
	} else {
		return file.Name
	}
}
type FileInfo struct {
	Name      string
	Path      string
	IsDir     bool
	Size      int64
	Mode      os.FileMode
	ModTime   int64
	User      string
	Group     string
	NumLinks  int
	BlockSize int64
}
func getFileDir(add string, showAll bool) ([]FileInfo, int64, bool, error) {
	// path, err := os.Getwd()
	// if add != "." {
	// 	path = path + "/" + add
	// }
	// if err != nil {
	// 	fmt.Printf("Error getting current directory: %v\n", err)
	// 	return nil, 0, err
	// }
	path, filename, v := givePath(add)
	var files []FileInfo
	var totalBlocks int64
	dir, err := os.Open(path)
	if err != nil {
		return nil, 0, v, err
	}
	defer dir.Close()
	entries, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("here we are!!!!")
		return nil, 0, v, err
	}
	for _, entry := range entries {
		fullPath := path + "/" + entry.Name()
		file, err := getSingleFileInfo(fullPath, entry)
		if err != nil {
			continue
		}
		if !showAll && file.Name[0] == '.' {
			totalBlocks += 0
		}else{
			totalBlocks += file.BlockSize
		}
		if filename == add && entry.Name() == filename {
			files = append(files, file)
			break
		}else if filename != add {
			files = append(files, file)
		}
	}
	return files, totalBlocks, v, nil
}
func givePath(add string) (path, filename string, print bool) {
	var finalpath string
	current, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}
	if add != "." {
		finalpath = current + "/" + add
	}else{
		finalpath = current
	}
	dir, err := os.Open(finalpath)
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)	
	}
	defer dir.Close()
	fileInfo, err := dir.Stat()
    if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1) 
	}
	if !fileInfo.IsDir() {
        return current, add, false
    }else{
		return finalpath, "", true
	}
}
func GetSingleFileInfo(path string) (FileInfo, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return FileInfo{}, err
	}
	return getSingleFileInfo(path, info)
}
func getSingleFileInfo(path string, info os.FileInfo) (FileInfo, error) {
	file := FileInfo{
		Name:    info.Name(),
		Path:    path,
		IsDir:   info.IsDir(),
		Size:    info.Size(),
		Mode:    info.Mode(),
		ModTime: info.ModTime().Unix(),
	}
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		file.NumLinks = int(stat.Nlink)
		file.BlockSize = stat.Blocks
		file.User = getUserName(stat.Uid)
		file.Group = getGroupName(stat.Gid)
		if file.IsDir {
			subdirs, err := countSubdirectories(path)
			if err == nil {
				file.NumLinks = subdirs + 2
			}
		}
	}
	return file, nil
}
func countSubdirectories(path string) (int, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}
	count := 0
	for _, entry := range entries {
		if entry.IsDir() {
			count++
		}
	}
	return count, nil
}
func getUserName(uid uint32) string {
	u, err := user.LookupId(strconv.FormatUint(uint64(uid), 10))
	if err != nil {
		return strconv.FormatUint(uint64(uid), 10)
	}
	return u.Username
}
func getGroupName(gid uint32) string {
	g, err := user.LookupGroupId(strconv.FormatUint(uint64(gid), 10))
	if err != nil {
		return strconv.FormatUint(uint64(gid), 10)
	}
	return g.Name
}
func formatTime(modTime int64) string {
	t := time.Unix(modTime, 0)
	now := time.Now()
	sixMonthsAgo := now.AddDate(0, -6, 0)
	if t.After(sixMonthsAgo) {
		return t.Format("Jan _2 15:04")
	}
	return t.Format("Jan _2  2006")
}