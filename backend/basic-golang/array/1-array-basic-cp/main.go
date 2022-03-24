package main

import "fmt"

// Buatlah array dengan tipe data string.
// Dimana size array adalah karakter dari setiap huruf nama panggilan kalian.
// gunakan inisialisasi menggunakan `:=` pada array1.
func main() {
	// TODO: answer here

	arr1 := [2]int8{10, 11}
	var arr2 [2]string

	arr2[0] = "satu"
	arr2[1] = "dua"

	fmt.Println(arr1)
	fmt.Println(arr2)

	// var txt = "Hello"

	// fmt.Printf("%s\n", txt)
	// fmt.Printf("%q\n", txt)
	// fmt.Printf("%8s\n", txt)
	// fmt.Printf("%-8s\n", txt)
	// fmt.Printf("%x\n", txt)
	// fmt.Printf("% x\n", txt)
}
