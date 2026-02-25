package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func gols(files []string, directories []string) error {
	var err error

	for _, f := range files {
		err = print_regular_file(f)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	if len(files) > 0 && len(directories) > 0 {
		fmt.Println()
	}

	for i, f := range directories {
		err = print_directory(f, len(directories) > 1)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if i < len(directories)-1 {
			fmt.Println()
		}
	}
	return err
}

func print_regular_file(file_name string) error {
	info, err := os.Lstat(file_name)
	if err != nil {
		return err
	}
	if info.Mode().IsRegular() && (info.Mode()&0111) != 0 {
		fmt.Print(Green)
	}
	fmt.Println(file_name, Reset)

	return nil
}

func print_directory(file_name string, many bool) error {
	var err error
	var list []string

	if many {
		fmt.Println(file_name + ":")
	}
	if file_name == "." {
		file_name, err = os.Getwd()
	}

	entries, err := os.ReadDir(file_name)
	if err != nil {
		return err
	}
	for _, f := range entries {
		if f.Name()[0] != '.' {
			list = append(list, f.Name())
		}
	}
	list = sort_directory(list)
	for _, f := range list {
		file_path := filepath.Join(file_name, f)
		info, err := os.Lstat(file_path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			fmt.Print(Blue)
		} else if info.Mode().IsRegular() && (info.Mode()&0111) != 0 {
			fmt.Print(Green)
		}
		fmt.Println(f, Reset)
	}

	return nil
}
