package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Dir struct {
	path    string
	files   []string
	subDirs []*Dir
}

func Build(path string) (*Dir, error) {
	d := &Dir{path: path}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			sd, err := Build(path + "/" + file.Name())
			if err != nil {
				fmt.Println(err)
				continue
			}
			if file.Name()[0] == '.' {
				continue
			}
			d.subDirs = append(d.subDirs, sd)
		} else {
			d.files = append(d.files, file.Name())
		}
	}

	return d, nil
}

func (d *Dir) Print(prefix string) {
	for i, fileName := range d.files {
		if i == len(d.files)-1 && len(d.subDirs) == 0 {
			fmt.Printf("%s└── %s\n", prefix, fileName)
		} else {
			fmt.Printf("%s├── %s\n", prefix, fileName)
		}
	}

	for i, dir := range d.subDirs {
		splitPath := strings.Split(dir.path, "/")
		lastFileName := splitPath[len(splitPath)-1]

		if i == len(d.subDirs)-1 {
			fmt.Printf("%s└── %s\n", prefix, lastFileName)
			dir.Print(prefix + "    ")
		} else {
			fmt.Printf("%s├── %s\n", prefix, lastFileName)
			dir.Print(prefix + "│   ")
		}
	}
}
