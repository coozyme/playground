package main

import (
	"fmt"
	"math"
)

// Check Point:
// Two Dimention Area
// - Input: Select Choice
// - Output: Result Operation

// Contoh:
// Input:
// 1: Rectange Area
// 2: Rectangular Area
// 3: Triangle Area
// 4: Circle Area
// - Enter choice 1, 2, 3, or 4: 1 | 2 | 3 | 4
//   - (1) Rectange Area
//   	- Masukkan sisi : 5
//   - (2) Rectangular Area
// 		- Masukkan panjang : 5
// 		- Masukkan lebar : 10
// 	 - (3) Triangle Area
// 		- Masukkan panjang alas segitiga: 5
// 		- Masukkan tinggi segitiga: 10
// 	 - (4) Circle Area
//      - Masukkan jari-jari : 4

// Output:
// - (1) Luas Persegi adalah : 25
// - (2) Luas Persegi Panjang adalah : 50
// - (3) Luas Segitiga adalah : 25
// - (4) Luas Lingkaran adalah : 50.240000

func main() {
	// TODO: answer here

	var InputArea int
	// fmt.Scan(InputArea)

	fmt.Println("Input : \n" +
		"1: Rectange Areas \n" +
		"2: Rectangular Area \n" +
		"3: Triangle Area \n" +
		"4: Circle Area \n")

	fmt.Print("Enter choice 1  2 | 3 | 4 : ")
	fmt.Scan(&InputArea)

	switch InputArea {
	case 1:
		var sisi, formula int64
		fmt.Println("==== Rectange Area ====")
		fmt.Print("Masukan sisi :")
		fmt.Scan(&sisi)

		formula = sisi * sisi

		fmt.Printf("Luas Persegi adalah %d \n", formula)

	case 2:
		var panjang, lebar, formula int64
		fmt.Println("==== Rectangular Area ====")
		fmt.Print("Masukan Panjang : ")
		fmt.Scan(&panjang)
		fmt.Print("Masukan Lebar : ")
		fmt.Scan(&lebar)

		formula = panjang * lebar

		fmt.Printf("Luas Persegi Panjang adalah %d \n", formula)
	case 3:
		var panjang, tinggi, formula int64
		fmt.Println("==== Triangle Area ====")
		fmt.Print("Masukkan panjang alas segitiga : ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan tinggi segitiga : ")
		fmt.Scan(&tinggi)

		formula = panjang * tinggi / 2

		fmt.Printf("Luas Segitiga adalah %d \n", formula)
	case 4:
		var jari, formula float64
		const pi = 3.14
		fmt.Println("==== Circle Area ====")
		fmt.Print("Masukkan jari - jari : ")
		fmt.Scan(&jari)

		formula = pi * (math.Pow(jari, 2))

		fmt.Printf("Luas Persegi Panjang adalah %f \n", formula)
	}

}
