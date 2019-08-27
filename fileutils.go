package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func isDirectory(path string) bool {
	fd, err := os.Stat(path)

	if err != nil {
		panic(err)
	}

	switch mode := fd.Mode(); {
	case mode.IsDir():
		return true
	case mode.IsRegular():
		return false
	}
	return false
}

func getFileList(dirPath string) []string {
	fileList := []string{}

	filepath.Walk(dirPath, func(path string, f os.FileInfo, err error) error {
		if debug {
			fmt.Println("PATH ==> " + path)
		}

		if isDirectory(path) {
			// Do nothing
			return nil
		}

		fileList = append(fileList, path)

		return nil
	})

	return fileList
}
