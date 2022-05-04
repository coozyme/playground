package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("dummyCommit")
}

// Gunakan struct untuk menyimpan data file nya ya
type FileData struct {
	Name string
	Size int
	Data []byte
}

func ReadFile(name string) (FileData, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatalf("failed read file, err: %s", err)
	}

	files := FileData{
		Name: name,
		Size: len(data),
		Data: data,
	}
	return files, nil // TODO: replace this
}
