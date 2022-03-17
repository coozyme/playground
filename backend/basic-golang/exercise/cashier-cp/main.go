package main

import (
	"fmt"
)

// Check Point:
// Cashier
// - Input: Nama Menu, Add Again
// - Output: Total Harga yang harus dibayar

// Contoh:
// Menu makanan:
// 1 . Menu: bakso ,  Price : 20000
// 2 . Menu: burger ,  Price : 15000
// 3 . Menu: sate ,  Price : 25000
// 4 . Menu: sosis ,  Price : 20000
// 5 . Menu: soto ,  Price : 25000

// Input:
// - Form:
//   - Masukan nama menu yang mau dipesan: bakso

// Output:
// Menu telah ditambahkan ke Ordered Menu:
// Menu: bakso ,  Price : 20000

// Input:
// - Ingin memesan menu lain?(yes/no): if no (break) --> if yes (add again)

// Output:
// Data Order Updated:
//  - Menu: sate ,  Price : 25000
//  - Menu: burger ,  Price : 15000
//  - Menu: sosis ,  Price : 20000
//  - Total harga makanan yang harus anda bayar:  60000

func main() {
	foodMenu := map[string]int64{
		"bakso":  20000,
		"burger": 15000,
		"sate":   25000,
		"sosis":  20000,
		"soto":   25000,
	}

	// TODO: answer here

	orderMenu := make(map[string]int64)

	for {
		fmt.Println("Menu makanan")
		for nama, harga := range foodMenu {
			fmt.Println("-", nama, ", Price: ", harga)
		}

		fmt.Println(" ")
		var NamaMenu, AddAgain string

		fmt.Println("Form :")
		fmt.Printf("- Masukkan nama menu yang mau dipesan :")
		fmt.Scan(&NamaMenu)

		if v, found := foodMenu[NamaMenu]; found {
			fmt.Println("Menu telah ditambahkan ke Ordered Menu:")
			fmt.Printf("Menu: %s, Price: %d \n", NamaMenu, v)
			orderMenu[NamaMenu] = v

		} else {
			fmt.Println("Menu Pesanan Anda Tidak Tersedia`")
		}

		fmt.Println()
		fmt.Printf("Ingin memesan menu lain? (yes/no): ")
		fmt.Scan(&AddAgain)
		if AddAgain == "no" {
			var total int64 = 0
			fmt.Println()
			fmt.Printf("Menu yang anda pesan : \n")
			for menu, price := range orderMenu {
				fmt.Printf("%s = %d \n", menu, price)
				total += price
			}
			fmt.Println("Total harga makanan yang harus anda bayar: ", total)
			break
		}

	}

}
