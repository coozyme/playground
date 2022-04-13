package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan looping pada template.
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untuk setiap user
// Pada setiap loop, cetak "Peringkat ke-n: Nama", contoh: "Peringkat ke-1: Roger"

type UserRank struct {
	Name  string
	Email string
	Rank  int
}

type Leaderboard struct {
	Users []*UserRank
}

func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
	var textTemplate string
	// TODO: answer here

	textTemplate = `{{range . }}Peringkat ke-{{ .Rank }}: {{ .Name }}{{end}}`

	t, err := template.New("3-loop-cp").Parse(textTemplate)

	if err != nil {
		panic(err)
	}

	var b bytes.Buffer

	err = t.Execute(&b, leaderboard.Users)

	if err != nil {
		panic(err)
	}

	return b.Bytes(), err

}

func main() {
	var leaderboardObject Leaderboard
	users := []*UserRank{
		{
			Name: "Roger",
			Rank: 1,
		},
		{
			Name: "Tony",
			Rank: 2,
		},
		{
			Name: "Bruce",
			Rank: 3,
		},
		{
			Name: "Natasha",
			Rank: 4,
		},
		{
			Name: "Clint",
			Rank: 5,
		},
	}

	leaderboardObject = Leaderboard{
		Users: users,
	}
	av, _ := ExecuteToByteBuffer(leaderboardObject)
	fmt.Println(string(av))
}
