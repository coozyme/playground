package main

import "fmt"

//buat struct Rectangle dengan dua atribut yaitu Width dan Length
// tambah dua method :
// SetWidthPointer(width int) dan SetWidthValue(width int)
// SetWidthPointer(width int) untuk mengubah width dengan pointer receiver
// SetWidthValue(width int) untuk mengubah width dengan value

// TODO: answer here
type Rectangle struct {
	Width  int
	Length int
}

func (r *Rectangle) SetWidthPointer(p int) {
	r.Width = p
	fmt.Println("in method SetWidthPointer", r.Width)
}

func (r Rectangle) SetWidthValue(p int) {
	r.Width = p
	fmt.Println("in method SetWidthValue", r.Width)
}

func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20

	fmt.Println("sebelum melakukan set width dengan pointer", r.Width)
	r.SetWidthPointer(30)
	fmt.Println("sesudah melakukan set width dengan pointer", r.Width)
	r.SetWidthValue(70)
	fmt.Println("sesudah melakukan set width dengan value", r.Width)
}
