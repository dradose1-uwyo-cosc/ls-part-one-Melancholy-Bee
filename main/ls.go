package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func gols(file_name string, is_dir bool, many_files bool, last_file bool) error {
	var err error
	var list []string

	if is_dir {
		if many_files {
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
	} else {
		info, err := os.Lstat(file_name)
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() && (info.Mode()&0111) != 0 {
			fmt.Print(Green)
		}
		fmt.Println(file_name, Reset)
	}

	if many_files && is_dir && !last_file {
		fmt.Println()
	}

	return err
}

func print_regular_files(files []string) {

}

func print_directories(files []string) {

}
