package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

//Buat struct Student dengan field Name tipe data <string> dan ScoreAverage tipe data <float64>
//tampilkan dengan wording:
//Hello Rogu,
//Nilai rata rata kamu 7.8

type Student struct {
	// TODO: answer here
	Name  string
	Nilai float32
}

// main function
func main() {
	buff := new(bytes.Buffer)
	// TODO: answer here
	std := Student{Name: "Rogu", Nilai: 7.8}
	tmp1 := template.New("Template_1")

	tmp1, err := tmp1.Parse("Hello {{.Name}},\nNilai rata rata kamu {{.Nilai}}")

	if err != nil {
		log.Fatalf("err parse %s", err.Error())
	}

	if err := tmp1.Execute(buff, std); err != nil {
		log.Fatalf("execute template error: %s", err.Error())
	}
	fmt.Println(buff.String())
}
