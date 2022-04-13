package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func main() {
	fmt.Print("dummy commit")
}

func CSVToMap(data map[string]string, fileName string) (map[string]string, error) {
	// TODO: answer here
	dataFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer dataFile.Close()

	r := csv.NewReader(dataFile)

	if _, err := r.Read(); err != nil {
		return  *new(map[string]string) ,err
	}

	dt, err := r.ReadAll()
	if err != nil {
		return  *new(map[string]string) ,err
	}

	return *new(map[string]string, nil
}
