//jangan ditunjukkan ke peserta
//mungkin dikerjakan jika waktu cukup aja
package main

import "fmt"

// fungsi ini digunakan untuk menambahkan point
// fungsi ini merupakan closure
func points(base int) func(x int) int {
	// TODO: answer here
	// basePoint := 0
	basePoint := base
	fmt.Println("a", basePoint)
	return func(x int) int {
		fmt.Println("ss", basePoint+x)
		basePoint = basePoint + x
		return basePoint

	}

}

func main() {
	toni := points(100) // base value 100
	tono := points(100)
	fmt.Println(toni(2))   // jadi 102
	fmt.Println(toni(3))   // 105
	fmt.Println(toni(100)) // 205
	fmt.Println(tono(2))   // 102
}
