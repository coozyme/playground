package main

import (
	"fmt"
)

type rect struct {
	len, width int
}

func (r rect) area() {
	fmt.Println(r.len * r.width)
}

func hello(num ...int) {
	num[0] = 22
	fmt.Println(num)
}

func main() {
	// s := make(map[string]int)
	// delete(s, "h")
	// fmt.Println(s["h"])

	// i := 2
	// s := "1000"
	// if len(s) > i {
	// 	i, _ := strconv.Atoi(s)
	// 	fmt.Println(i)
	// 	i = i + 5
	// 	fmt.Println(i)
	// }
	// fmt.Println("sop", i)

	// r := &rect{len: 5, width: 6}
	// r.area()

	// i := []int{5, 6, 7}
	// hello(i...)
	// fmt.Println(i[0])

	// keyboard := []int32{40, 50, 60}
	// drivers := []int32{5, 8, 12}
	// b := 60
	// max := int32(-1)
	// for i := int32(0); i < int32(len(keyboard)); i++ {
	// 	for j := int32(0); j < int32(len(drivers)); j++ {
	// 		sum := keyboard[i] + drivers[i]
	// 		if sum > max && sum <= int32(b) {
	// 			max = sum
	// 			fmt.Println("more 1", max)
	// 		} else if sum < max && sum >= int32(b) {
	// 			if max == int32(b) {

	// 				fmt.Println("more 2", max)
	// 			}

	// 		}
	// 	}
	// }

	var i, j, k, l int

	for i = 11; i > 0; i-- {
		for j = 1; j <= i; j++ {
			fmt.Printf(" ")
			// fmt.Printf("\n")
		}
		for k = 10; k > i; k-- {
			fmt.Printf("*")
		}
		for l = 11; l > i; l-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

}
