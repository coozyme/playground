package main

import "fmt"

//fungsi square akan mengembalikan nilai pangkat 2
//dari dua parameter yang diterima
//contoh
//parameter yang diterima (4,3)
//maka akan mengembalikan 16 dan 9
func main() {

	result1, result2 := square(4, 5)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(square(9, 8))
}

//gunakan * untuk melakukan perkalian
// TODO: answer here

func square(num1, num2 int) (int, int) {
	count1 := num1 * num1
	count2 := num2 * num2

	return count1, count2
}
