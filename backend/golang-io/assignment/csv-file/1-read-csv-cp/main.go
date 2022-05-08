package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("dummy commit")

	data := make(map[string]string)
	testData, _ := CSVToMap(data, "questions.csv")
	// fmt.Println("testData", testData["Nama ibukota indonesia?"])
	fmt.Println("testData", testData)
}

func CSVToMap(data map[string]string, fileName string) (map[string]string, error) {
	// TODO: answer here

	dataFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer dataFile.Close()

	r := csv.NewReader(dataFile)

	if _, err = r.Read(); err != nil {
		return *&data, err
	}

	makeNote := make(map[string]string)

	listData, err := r.ReadAll()
	if err != nil {
		return *&data, err
	}

	for _, v := range listData {
		fmt.Println("kj", v[0], v[1])
		makeNote[strings.TrimSpace(v[0])] = strings.TrimSpace(v[1])
	}

	return *&makeNote, nil
}
