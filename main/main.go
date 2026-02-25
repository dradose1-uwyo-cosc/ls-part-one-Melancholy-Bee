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

	if len(os.Args) == 1 {
		files = []string{"."}
	} else {
		files = os.Args[1:]
	}
	files, directories, err = sort_many_files(files)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	err = gols(files, directories)

	if err != nil {
		os.Exit(1)
	}
}
