package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Dari contoh sebelumnya tambahkan filter untuk
// menampilkan meja berdasarkan total meja
// key input dari client menggunakan `total`
// contuh URL: http://localhost:8080/table?total=2

type Table struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Color string `json:"color"`
	Total int    `json:"total"`
}

func TableHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {

		// TODO: answer here
		fieldData := []Table{}
		total := r.FormValue("total")
		convTotal, _ := strconv.Atoi(total) //convert from string -> int

		if total == "" {
			http.Error(w, "invalid total", http.StatusBadRequest)
			return
		}

		for _, t := range data {
			if t.Total == convTotal {
				fieldData = append(fieldData, t)
			}
		}

		result, _ := json.Marshal(fieldData)

		if len(fieldData) == 0 {
			http.Error(w, `{"status":"table not found"}`, http.StatusNotFound)
			return
		}

		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

var data = []Table{
	{ID: "T001", Type: "Meja Lipat", Color: "Coklat", Total: 3},
	{ID: "T002", Type: "Meja Belajar", Color: "Hitam", Total: 4},
	{ID: "T003", Type: "Meja Makan", Color: "Coklat", Total: 9},
	{ID: "T004", Type: "Meja Hejau", Color: "Hijau", Total: 3},
}

func main() {
	http.HandleFunc("/table", TableHandler)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
