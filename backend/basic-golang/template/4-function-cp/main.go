package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan function pada template.
// Lengkapi function CalculateScore yang berfungsi untuk menjumlahkan total score dari users
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untuk setiap user
// Pada setiap loop, cetak "Nama: Score", contoh: "Roger: 1000"
// Pada bagian terakhir, cetak "Total Score: total", contoh: "Total Score: 50000"

type UserRank struct {
	Name  string
	Email string
	Rank  int
	Score int
}

type Leaderboard struct {
	Users []*UserRank
}

func CalculateScore(leaderboard Leaderboard) int {
	// TODO: answer here
	scr := 0
	// scr += leaderboard.Users÷
	for _, v := range leaderboard.Users {
		fmt.Println("po", v.Score)
		scr += v.Score
	}
	return scr
}

func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
	var textTemplate string
	// TODO: answer here

	// score := CalculateScore(leaderboard)

	funcTotalScore := template.FuncMap{
		"sum": CalculateScore(leaderboard),
	}

	textTemplate = "{{range .}} {{ .Name }}: {{ .Score }} {{end}} {{ .scrore }}"

	t, errs := template.New("exec").Funcs(funcTotalScore).Parse(textTemplate)
	// t, errs := template.New("exec").Parse(textTemplate)

	if errs != nil {
		panic(errs)
	}

	var b bytes.Buffer

	// fmt.Println(nm)

	errs = t.Execute(&b, leaderboard.Users)
	return b.Bytes(), errs
}

func main() {
	ary := []*UserRank{
		{
			Name:  "Ary",
			Email: "ary@gmail.com",
			Rank:  2,
			Score: 5,
		},
		{
			Name:  "Ary Setya",
			Email: "arysetya@gmail.com",
			Rank:  22,
			Score: 53,
		},
	}

	abc := Leaderboard{Users: ary}

	ab, er := ExecuteToByteBuffer(abc)
	if er != nil {
		fmt.Println("err:", er)
	}
	fmt.Println("ok", string(ab))

}
