package main

import (
	"fmt"
	"math"
)

// Check Point:
// Menghitung volume tabung
// - Input: jari-jari, tinggi
// - Output: volume tabung

// Contoh:
// Input:
// - Masukkan jari-jari alas tabung: 2
// - Masukkan tinggi tabung : 20
// Output:
// - Jadi volumenya adalah : 251.200012

func main() {
	// TODO: answer here

	var InputJari, InputTinggi, formula float64

	fmt.Printf("Masukan jari-jari alas tabung: ")
	fmt.Scan(&InputJari)
	fmt.Printf("Masukan Tinggi alas tabung: ")
	fmt.Scan(&InputTinggi)

	const pi = 3.14
	formula = pi * math.Pow(InputJari, 2) * InputTinggi

	fmt.Printf("jari = %f dan tinggi %f pi = %f formula = %f \n", math.Pow(InputJari, 2), InputTinggi, pi, formula)
	fmt.Printf("Jadi volumenya adalah : %f \n", formula)

}
