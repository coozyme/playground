package main

import (
	"encoding/json"
	"fmt"
)

// Dari struct dan nama field yang sama dari contoh
// Buat string JSON dengan hasil seperti berikut
// {"jenis":"Meja Belajar","color":"green","jumlah":2}

type Meja struct {
	// TODO: answer here
	Jenis  string `json:"jenis"`
	Warna  string `json:"color"`
	Jumlah int    `json:"jumlah"`
}

func (m Meja) EncodeJSON() string {
	// TODO: answer here
	mr, _ := json.Marshal(m)
	return string(mr)
}

func NewMeja(m Meja) Meja {
	return m
}

func main() {
	m := Meja{
		Jumlah: 3,
		Jenis:  "ok",
		Warna:  "merah",
	}

	fmt.Println(m.EncodeJSON())
}
