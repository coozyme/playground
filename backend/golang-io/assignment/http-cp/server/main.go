package main

import (
	"net/http"
	"strconv"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		w.Write([]byte("Hello, world!"))
	})
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		params := r.FormValue("data")

		w.Write([]byte(params))
	})
	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		var b []byte
		// params1 := r.FormValue("a")
		// params2 := r.FormValue("b")

		// convertParams1, _ := strconv.Atoi(params1)
		// convertParams2, _ := strconv.Atoi(params2)
		params := r.URL.Query()["a","b"]

		sum := convertParams1 + convertParams2
		b[0] = byte(sum)

		w.Write([]byte(b))
	})
	mux.HandleFunc("/hellojson", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
	})

	return mux
}
func main() {
	http.ListenAndServe(":8080", Routes())
}
