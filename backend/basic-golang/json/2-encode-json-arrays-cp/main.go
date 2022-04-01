package main

import (
	"encoding/json"
	"fmt"
)

// Buat string JSON dengan hasil seperti berikut
// [{"jenis":"Meja Lipat","warna":"Coklat","jumlah":40,"deskripsi":"meja untuk belajar"},{"jenis":"Meja Hijau","warna":"Hijau","jumlah":10,"deskripsi":"meja untuk pengadilan"}]

type Meja struct {
	// TODO: answer here
	Jenis     string `json:"jenis"`
	Warna     string `json:"warna"`
	Jumlah    int    `json:"jumlah"`
	Deskripsi string `json:"deskripsi"`
}

type Items struct {
	MejaMeja []Meja
}

func (m Items) EncodeJSON() string {
	// TODO: answer here
	encode, e := json.Marshal(m.MejaMeja)
	if e != nil {
		return "err"
	}
	return string(encode)
}

func NewMeja(m Items) Items {
	return m
}

func main() {
	item := []Meja{{
		Jenis:     "Meja Lipat",
		Warna:     "Coklat",
		Jumlah:    40,
		Deskripsi: "meja untuk belajar",
	},
		{
			Jenis:     "Meja Hijau",
			Warna:     "Hijau",
			Jumlah:    10,
			Deskripsi: "meja untuk pengadilan",
		},
	}

	fmt.Println(item)

}
