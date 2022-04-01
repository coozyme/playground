package main

import (
	"encoding/json"
	"fmt"
)

// Dari contoh yang telah diberikan, cobalah untuk melakukan decode dari json menjadi objek dari struct.
// Mengambil kasus pada encode, lengkapi function DecodeToLeaderboard agar json dapat di decode menjadi objek Leaderboard

// type UserRank struct {
// 	Name  string
// 	Email string
// 	Rank  int
// }

type UserRank struct {
	// TODO: answer here
	Name  string `json:"name"`
	Email string `json:"email.omitempty"`
	Rank  int    `json:"rank"`
}

type Leaderboard struct {
	Users []*UserRank
}

func DecodeToLeaderboard(jsonData []byte) (Leaderboard, error) {
	// TODO: answer here

	u := Leaderboard{}

	er := json.Unmarshal(jsonData, &u)
	if er != nil {
		return u, er

	}
	return u, nil

}

func main() {
	a, e := DecodeToLeaderboard([]byte(`{
				"Users":[
				   {
					  "Name":"Roger",
					  "Email":"roger@ruangguru.com",
					  "Rank":1
				   },
				   {
					  "Name":"Tony",
					  "Email":"tony@ruangguru.com",
					  "Rank":2
				   },
				   {
					  "Name":"Bruce",
					  "Email":"bruce@ruangguru.com",
					  "Rank":3
				   },
				   {
					  "Name":"Natasha",
					  "Email":"natasha@ruangguru.com",
					  "Rank":4
				   },
				   {
					  "Name":"Clint",
					  "Email":"clint@ruangguru.com",
					  "Rank":5
				   }
				]
			 }`))

	fmt.Println(a)
	fmt.Println(e)

}
