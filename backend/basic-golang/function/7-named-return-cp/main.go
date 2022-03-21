package main

import "fmt"

//fungsi square sama seperti sebelumnya
//yang membedakan adalah fungsi ini
//menggunakan named return
func main() {
	result1, result2 := square(4, 5)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(square(9, 8))
}

// TODO: answer here
func square(num1, num2 int) (count1, count2 int) {
	count1 = num1 * num1
	count2 = num2 * num2

	return count1, count2
}
