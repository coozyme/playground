package main

import "fmt"

//pakai * untuk melakukan perkalian
//misal number1 * number2
func main() {
	result := multiply(3, 3)
	fmt.Println(result)
	fmt.Println(multiply(5, 5))
}

// TODO: answer here

func multiply(val1, val2 int) int {
	return val1 * val2
}
