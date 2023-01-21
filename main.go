package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	d, _ := Build(".")
	fmt.Println(".")
	d.Print("")
}

func printTree(path string, prefix string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, file := range files {
		if file.IsDir() {
			if i == len(files)-1 {
				fmt.Printf("%s└── %s\n", prefix, file.Name())
				printTree(filepath.Join(path, file.Name()), prefix+"    ")
			} else {
				fmt.Printf("%s├── %s\n", prefix, file.Name())
				printTree(filepath.Join(path, file.Name()), prefix+"│   ")
			}
		} else {
			if i == len(files)-1 {
				fmt.Printf("%s└── %s\n", prefix, file.Name())
			} else {
				fmt.Printf("%s├── %s\n", prefix, file.Name())
			}
		}
	}
}
