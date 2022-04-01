package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan empty interface.
// Buatlah dua data makanan dan minuman yaitu ayam goreng dan cola yang memiliki atribut
// Nama, Jenis, dan Harga.
// Ayam Goreng, Cepat saji, 20000
// Cola, Minuman, 7000

func GetMenu() []map[string]interface{} {
	var menu []map[string]interface{}

	// TODO: answer here
	ayam := make(map[string]interface{})
	ayam["Nama"] = "Ayam Goreng"
	ayam["Jenis"] = "Cepat saji"
	ayam["Harga"] = 20000

	cola := make(map[string]interface{})
	cola["Nama"] = "Cola"
	cola["Jenis"] = "Minuman"
	cola["Harga"] = 7000

	menu = append(menu, ayam, cola)

	return menu
}

func main() {
	menu := GetMenu()

	for _, m := range menu {
		for k, v := range m {
			fmt.Printf("%s: %v\n", k, v)
		}
	}
}
