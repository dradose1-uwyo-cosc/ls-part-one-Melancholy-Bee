package functions

import (
	"errors"
	"os"
	"sort"
	"strings"
)

// Sort the arguments by files and directories
func Sort_many_files(files []string) ([]string, []string, error) {
	var f []string
	var dir []string
	var errs []error

	for _, file := range files {
		info, err := os.Lstat(file)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if info.IsDir() {
			dir = append(dir, file)
		} else {
			f = append(f, file)
		}
	}
	sort.Strings(f)
	sort.Strings(dir)

	return f, dir, errors.Join(errs...)
}

// Sort the files in the specified directory
func sort_directory(files []string) []string {
	// Apply non-case sensitive sorting while ignoring periods like ls
	sort.Slice(files, func(i, j int) bool {
		a := strings.ReplaceAll(strings.ToLower(files[i]), ".", "")
		b := strings.ReplaceAll(strings.ToLower(files[j]), ".", "")
		return a < b
	})
	return files
}
