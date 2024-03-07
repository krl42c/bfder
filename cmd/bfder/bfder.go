package bfder

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func Run(search string, path string) {
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	bsearch := []byte(search)
	files, err := read_dir(path)
	if err != nil {
		fmt.Println("No files found")
	}

	for _, f := range files {
		f = path + f
		data, err := read_file(f)
		if err != nil {
			fmt.Println(err)
		} else {
			if bytes.Contains(data, bsearch) {
				in := bytes.Index(data, bsearch)
				fmt.Printf("{%s}: %v\n", f, in)
			}
		}
	}
}

func read_dir(path string) ([]string, error) {
	var files []string
	fileInfo, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range fileInfo {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}

	return files, nil
}

func read_file(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}
