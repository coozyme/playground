package main

import (
	"log"
	"os"
)

// dalam test ini terdapat fungsi os.Remove ya. itu automatis nge remove file yang telah dibuat
// Untuk keperluan testing
func WriteFile(fileName string, fileData string) error {
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("failed create file, err: %s", err)
	}

	defer file.Close()

	len, err := file.WriteString(fileData)

	if err != nil {
		return err
	}

	log.Printf("len: %d", len)
	return nil // TODO: replace this
}
