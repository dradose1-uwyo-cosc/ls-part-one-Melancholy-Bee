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
	var reg_num int
	var many_files bool
	var last_file bool = false
	var err error

	if len(os.Args) == 1 {
		files = []string{"."}
	} else {
		files = os.Args[1:]
	}
	if len(files) == 1 {
		many_files = false
	} else {
		many_files = true
		files, reg_num, err = sort_many_files(files)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	for i, f := range files {
		if i == len(files)-1 {
			last_file = true
		}
		err = gols(f, i >= reg_num, many_files, last_file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
