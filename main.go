package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var src = ""

func main() {
	Em(src)
}

func Em(path string) {
	listDir, _ := ioutil.ReadDir(path)
	for _, file := range listDir {
		if file.IsDir() {
			Em(filepath.Join(path, file.Name()))
		}
	}

	if filepath.Base(filepath.Dir(path)) == filepath.Base(path) {

		log.Println("pp", filepath.Dir(path))
		move(path, filepath.Dir(path))
	}
}

func move(oldPath, newPath string) {
	var op, np string
	var err error
	listDir, _ := ioutil.ReadDir(oldPath)
	for _, file := range listDir {
		op = filepath.Join(oldPath, file.Name())
		np = filepath.Join(newPath, file.Name())
		if _, err := os.Stat(newPath); os.IsNotExist(err) && file.IsDir() {
			move(op, np)
		}
		err = os.Rename(op, np)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = os.Remove(oldPath)
	if err != nil {
		log.Fatal(err)
	}
}