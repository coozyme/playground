package main

import "fmt"

//memanggil fungsi goodMorning
//fungsi goodMorning akan melakukan print selamat pagi + name yang didapat dari parameter fungsi
func main() {
	goodMorning("teman")
	goodMorning("teman 2")

}

// TODO: answer here
func goodMorning(value string) {
	fmt.Println("hallo ", value)
}
