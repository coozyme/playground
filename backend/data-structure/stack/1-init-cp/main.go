package main

// Dari contoh yang telah diberikan, cobalah untuk membuat stack yang memiliki jumlah maksimal elemen sebanyak 10 dengan menambahkan atribut Size.
// Gunakan slice untuk menyimpan data pada stack. Pada kasus ini, data yang disimpan berupa int.
// Lengkapi function NewStack sehingga function tersebut dapat mengembalikan objek Stack dengan memiliki ukuran dari parameter.
type Stack struct {
	// TODO: answer here
	Top  int
	Data []int
	Size int
}

func NewStack(size int) Stack {
	// TODO: answer here
	return Stack{
		Top:  -1,
		Data: []int{},
		Size: size,
	}
}
