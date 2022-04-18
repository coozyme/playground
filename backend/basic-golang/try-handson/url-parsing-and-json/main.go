package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type Barang struct {
	Nama        string `json:"nama"`
	Harga       int    `json:"harga"`
	JenisBarang string `json:"jenis_barang"`
}

func main() {
	a := 4.85
	fmt.Println(a)
	fmt.Println(int(a))

	b := "900"
	c, _ := strconv.Atoi(b)
	fmt.Println(c + 400)

	var d = "http://www.belajar-golang.com/basic?name=Martha&tinggi=140"
	fmt.Println(d)

	e, _ := url.Parse(d)

	fmt.Println(e.Scheme)
	fmt.Println(e.Host)
	fmt.Println(e.Path)
	fmt.Println(e.Query())

	name := e.Query()["name"][0]
	t := e.Query()["tinggi"][0]
	fmt.Println(name)
	fmt.Println(t)
	DecodeJSONToObject()
	DecodeArrayJSONToObject()
	DecodeJSONToMapStringInterface()
	DecodeJSONToInterface()

	EncodeObjectToJSON()
}

func DecodeJSONToObject() {
	jsonString := `{"nama": "AAA", "harga": 78000, "jenis_barang": "Makanan"}`
	jsonData := []byte(jsonString)

	var Bar Barang

	err := json.Unmarshal(jsonData, &Bar)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Bar.Nama)
	fmt.Println(Bar.Harga)
	fmt.Println(Bar.JenisBarang)
}

func DecodeArrayJSONToObject() {
	jsonString := `[
		{"nama": "AAA", "harga": 78000, "jenis_barang": "Makanan"},
		{"nama": "BBB", "harga": 88000, "jenis_barang": "Minuman"},
		{"nama": "CCC", "harga": 98000, "jenis_barang": "Makanan"},
		{"nama": "DDD", "harga": 108000, "jenis_barang": "Minuman"}
	]`
	jsonData := []byte(jsonString)

	var Bar []Barang

	err := json.Unmarshal(jsonData, &Bar)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range Bar {
		fmt.Println("Nama barang", v.Nama)
		fmt.Println("Harga barang", v.Harga)
	}
}

func DecodeJSONToMapStringInterface() {
	jsonString := `{"nama": "AAA", "harga": 78000, "jenis_barang": "Makanan"}`
	jsonData := []byte(jsonString)

	var Bar map[string]interface{}

	err := json.Unmarshal(jsonData, &Bar)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Bar["jenis_barang"])
	fmt.Println(Bar["harga"])
}

func DecodeJSONToInterface() {
	jsonString := `{"nama": "AAA", "harga": 78000, "jenis_barang": "Makanan"}`
	jsonData := []byte(jsonString)

	var Bar interface{}

	err := json.Unmarshal(jsonData, &Bar)
	if err != nil {
		fmt.Println(err)
	}

	decodeBar := Bar.(map[string]interface{})

	fmt.Println(decodeBar["nama"])
}

func EncodeObjectToJSON() {
	data := Barang{"AAA", 78000, "Makanan"}
	res, _ := json.Marshal(data)

	fmt.Println(string(res))
}
