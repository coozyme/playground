package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3}

	res := SumAll(arr)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := SumAllCorrect(arr)
	// fmt.Println(resCorrect)
	SumAllCorrect([]int{12, 5, 10, 1, 3})
}

func SumAll(arr []int) int {
	res := 0
	for val := range arr {
		res += val
	}
	return res
}

func SumAllCorrect(arr []int) int {
	// TODO: replace this
	res := 0

	for _, val := range arr {
		fmt.Println("es", val)
		res += val
	}

	return res
}
