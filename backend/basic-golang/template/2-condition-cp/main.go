package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan condition pada template.
// Lengkapi function ExecuteToByteBuffer dan variabel textTemplate sehingga memiliki keluaran:
// 1. Jika saldo sama dengan 0, cetak "Akun Tony dengan Nomor 1002321 tidak memiliki saldo."
// 2. Jika saldo lebih dari 0, cetak "Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000."

type Account struct {
	Name    string
	Number  string
	Balance int
}

func ExecuteToByteBuffer(account Account) ([]byte, error) {
	var textTemplate string
	// TODO: answer here
	textTemplate = "{{if (gt .Balance 0) }} Akun {{ .Name }} dengan Nomor {{ .Number }} memiliki saldo sebesar ${{ .Balance }}.{{ else }} Akun {{ .Name }} dengan Nomor {{ .Number }} tidak memiliki saldo. {{ end }}"

	t, errs := template.New("2-condition-cp").Parse(textTemplate)

	if errs != nil {
		panic(errs)
	}
	var b bytes.Buffer

	errs = t.Execute(&b, account)

	if errs != nil {
		panic(errs)
	}

	return b.Bytes(), errs
}

func main() {
	account := Account{
		Name:    "Tony",
		Number:  "1002321",
		Balance: 1000,
	}
	b, err := ExecuteToByteBuffer(account)

	if err != nil {
		log.Printf("err@1 %s", err)
	}

	fmt.Println(string(b))
}
