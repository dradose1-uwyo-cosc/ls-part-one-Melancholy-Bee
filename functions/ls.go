package functions

import (
	"fmt"
	"os"
	"path/filepath"
)

type color string

const (
	Reset color = "\033[0m"
	Green color = "\033[01;32m"
	Blue  color = "\033[01;34m"
)

func Gols(files []string, directories []string) error {
	var err error
	colors := isTerminal(os.Stdout)

	// Print regular files
	for _, f := range files {
		err = print_regular_file(f, colors)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	// Add blank line after section if multiple argument types
	if len(files) > 0 && len(directories) > 0 {
		fmt.Println()
	}

	// Print directory contents
	for i, f := range directories {
		err = print_directory(f, len(directories)+len(files) > 1, colors)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if i < len(directories)-1 {
			fmt.Println()
		}
	}
	return err
}

func print_regular_file(file_name string, colors bool) error {
	info, err := os.Lstat(file_name)
	if err != nil {
		return err
	}
	if colors && info.Mode().IsRegular() && (info.Mode()&0111) != 0 {
		fmt.Print(Green)
	}
	fmt.Println(file_name, Reset)

	return nil
}

func print_directory(file_name string, many bool, colors bool) error {
	var err error
	var list []string

	// Add header if multiple arguments
	if many {
		fmt.Println(file_name + ":")
	}
	if file_name == "." {
		file_name, err = os.Getwd()
	}

	// Read in directory contents
	entries, err := os.ReadDir(file_name)
	if err != nil {
		return err
	}
	for _, f := range entries {
		// Avoid hidden files
		if f.Name()[0] != '.' {
			list = append(list, f.Name())
		}
	}
	// Sort directory contents lexically
	list = sort_directory(list)
	for _, f := range list {
		file_path := filepath.Join(file_name, f)
		info, err := os.Lstat(file_path)
		if err != nil {
			return err
		}
		if colors && info.IsDir() {
			fmt.Print(Blue)
		} else if colors && info.Mode().IsRegular() && (info.Mode()&0111) != 0 {
			fmt.Print(Green)
		}
		fmt.Println(f, Reset)
	}

	return nil
}
