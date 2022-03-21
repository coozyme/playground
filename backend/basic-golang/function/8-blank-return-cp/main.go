package main

import (
	"fmt"
	"math"
)

//fungsi square sama seperti sebelumnya
//yang membedakan adalah fungsi ini
//menggunakan blank return
func main() {
	result1, result2 := square(4, 5)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(square(9, 8))
}

// TODO: answer here
func square(num1, num2 int) (count1, count2 int) {
	count1 = int(math.Pow(float64(num1), 2))
	count2 = int(math.Pow(float64(num2), 2))

	return
}
