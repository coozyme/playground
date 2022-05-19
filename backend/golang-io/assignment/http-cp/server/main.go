package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type response struct {
	Message string `json:"message"`
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	})
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		params := r.FormValue("data")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(params))
	})
	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here

		queryOne := r.URL.Query()["a"]
		queryTwo := r.URL.Query()["b"]

		convertParams1, _ := strconv.Atoi(queryOne[0])
		convertParams2, _ := strconv.Atoi(queryTwo[0])

		sum := convertParams1 + convertParams2

		lk := strconv.FormatInt(int64(sum), 10)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(lk))
	})
	mux.HandleFunc("/hellojson", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		resp := make(map[string]string)
		resp["message"] = "Hello, world!"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		marshal, err := json.Marshal(resp)
		if err != nil {
			log.Printf("failed Encode, err: %v", err)
		}
		w.Write(marshal)
	})

	return mux
}
func main() {
	http.ListenAndServe(":8080", Routes())
}
