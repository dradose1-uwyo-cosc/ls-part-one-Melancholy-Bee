package functions

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Gols(files []string, directories []string, writer *bufio.Writer) error {
	var err error
	colors := isTerminal(os.Stdout)

	// Print regular files
	for _, f := range files {
		err = print_regular_file(f, colors, writer)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	// Add blank line after section if multiple argument types
	if len(files) > 0 && len(directories) > 0 {
		writer.WriteByte('\n')
	}

	// Print directory contents
	for i, f := range directories {
		err = print_directory(f, len(directories)+len(files) > 1, colors, writer)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if i < len(directories)-1 {
			writer.WriteByte('\n')
		}
	}
	return err
}

func print_regular_file(file_name string, colors bool, w *bufio.Writer) error {
	info, err := os.Lstat(file_name)
	if err != nil {
		return err
	}

	// Check if file is an executable
	if colors && info.Mode().IsRegular() && (info.Mode()&0111) != 0 {
		write_color(Green, w)
		io.WriteString(w, file_name)
		reset_color(w)
	} else {
		io.WriteString(w, file_name)
	}
	w.WriteByte('\n')

	return nil
}

func print_directory(file_name string, many bool, colors bool, w *bufio.Writer) error {
	var err error
	var list []string

	// Add header if multiple arguments
	if many {
		io.WriteString(w, file_name+":\n")
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
		if colors {
			if info.IsDir() {
				write_color(Blue, w)
			} else if info.Mode().IsRegular() && (info.Mode()&0111) != 0 {
				write_color(Green, w)
			}
			io.WriteString(w, f)
			reset_color(w)
		} else {
			io.WriteString(w, f)
		}
		w.WriteByte('\n')
	}

	return nil
}
