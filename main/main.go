package main

import (
	"fmt"
	"os"
)

type color string

const (
	Reset color = "\033[0m"
	Green color = "\033[01;32m"
	Blue  color = "\033[01;34m"
)

func main() {
	var files []string
	var directories []string
	var err error

	// Separate file arguments from program name
	if len(os.Args) == 1 {
		files = []string{"."}
	} else {
		files = os.Args[1:]
	}

	// Sort files and directories separately and lexically
	files, directories, err = sort_many_files(files)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Call ls command for printing
	err = gols(files, directories)

	// Exist abnormally if any errors occur
	if err != nil {
		os.Exit(1)
	}
}
