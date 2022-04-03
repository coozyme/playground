package main

import (
	"encoding/json"
	"fmt"
)

// buat JSON string array nested seperti berikut
/*
{
	"ruangTamu": {
			"items": [
					{
							"nama": "Meja",
							"jumlah": 20,
							"warna": "Coklat",
							"ukuran": {
									"panjang": "50 cm",
									"tinggi": "25 cm"
							}
					},
					{
							"nama": "Kursi",
							"jumlah": 1,
							"warna": "Hitam",
							"ukuran": {
									"panjang": "70 cm",
									"tinggi": "30 cm"
							}
					}
			]
	}
}
*/

// TODO: answer here
type Ukuran struct {
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
}
type Item struct {
	Nama   string `json:"nama"`
	Jumlah int    `json:"jumlah"`
	Warna  string `json:"warna"`
	Ukuran Ukuran `json:"ukuran"`
}

type RuangItems struct {
	RuangItems []Item `json:"items"`
}
type Ruang struct {
	RuangTamu RuangItems `json:"ruangTamu"`
}

func (r Ruang) EncodeJSON() string {
	// TODO: answer here
	encode, err := json.Marshal(r)
	if err != nil {
		return "JSON Marshal error: "
	}

	return string(encode)
}

func NewRuang(r Ruang) Ruang {
	return r
}

func main() {
	items := Ruang{
		RuangTamu: RuangItems{
			RuangItems: []Item{
				{
					Nama:   "Meja",
					Jumlah: 20,
					Warna:  "Coklat",
					Ukuran: Ukuran{
						Panjang: "50 cm",
						Tinggi:  "25 cm",
					},
				},
				{
					Nama:   "Kursi",
					Jumlah: 1,
					Warna:  "Hitam",
					Ukuran: Ukuran{
						Panjang: "70 cm",
						Tinggi:  "30 cm",
					},
				},
			},
		},
	}

	meja := NewRuang(items)
	result := meja.EncodeJSON()
	fmt.Println(result)

}
