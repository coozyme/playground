package main

import (
	"net/http"
)

var names = []string{
	"Tony",
	"Roger",
	"Bruce",
	"Stephen",
}

func IsNameExists(name string) bool {
	for _, n := range names {
		if n == name {
			return true
		}
	}

	return false
}

func GetNameHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		if r.Method == "GET" {
			name := r.URL.Query()["name"]

			if len(name) >= 1 {
				result := IsNameExists(name[0])
				if result {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
				}
			}

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/name", GetNameHandler())
	return mux
}
