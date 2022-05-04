package main

import "fmt"

// Disini kita akan mencoba untuk melakukan subslicing pada slice.
// Coba langsung gunakan function append ketika melakukan subslicing.
// contoh slice = append (slice, slice2[0:3])

// Silahkan copy slice dan mempunyai value "Marcus", "is", "known", "to", "be" dan "a", "philosopher"
// Silahkan print slice tersebut
// Expected output : [Marcus is known to be a philosopher]
func main() {
	slice := []string{"Marcus", "is", "known", "to", "be", "one", "of", "five", "greatest", "emperors", "of", "rome",
		"Aurelius", "is", "also", "known", "to", "be", "a", "philosopher"}

	// TODO: answer here
	var sliceCopy []string
	sliceCopy = append(sliceCopy, slice...)

	slice1 := sliceCopy[0:5]
	slice2 := sliceCopy[18:20]

	fmt.Println(append(slice1, slice2...))
	fmt.Println(append(sliceCopy[0:5], sliceCopy[18:20]...))
}
