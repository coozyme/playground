package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	// encrypt token dari username menggunakan bcrypt kemudian kirim ke user kedalam cookie
	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		cookieName := "token"
		var creds Credentials

		// Task:  Buat JSON body diconvert menjadi credential struct & return bad request ketika terjadi kesalahan decoding json

		// TODO: answer here
		err := json.NewDecoder(r.Body).Decode(&creds)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Task: Ambil password dari username yang dipakai untuk login

		// TODO: answer here
		expectPassword := creds.Password

		// Task: return unauthorized jika password salah

		// TODO: answer here

		checkUsername := expectPassword == users[creds.Username]
		if !checkUsername {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Task: 1. Buat token string menggunakan bcrypt dari credential username
		// 		 2. return internal server error ketika terjadi kesalahan encrypting token

		// TODO: answer here
		// token :=

		// Task: Set token string kedalam cookie response

		// TODO: answer here
	})

	return mux
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
