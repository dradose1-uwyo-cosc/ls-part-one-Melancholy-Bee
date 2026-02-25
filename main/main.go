package main

import (
	"bufio"
	"fmt"
	"gols/functions"
	"os"
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
	files, directories, err = functions.Sort_entries(files)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Call ls command for printing
	writer := bufio.NewWriter(os.Stdout)
	err = functions.Gols(files, directories, writer)
	defer writer.Flush()

	// Exist abnormally if any errors occur
	if err != nil {
		os.Exit(1)
	}
}
