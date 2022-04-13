package main

import (
	"fmt"
	"sort"
)

type Score struct {
	Name    string
	Correct int
	Wrong   int
	Empty   int
}
type Scores []Score

func (s Scores) Len() int {
	return len(s)
}

func (s Scores) Less(i, j int) bool {
	aTotal := s[i].Correct*4 - s[i].Wrong // Total perhitungan index i
	bTotal := s[j].Correct*4 - s[j].Wrong // Total perhitungan index j

	// Jika totalnya sama
	if aTotal == bTotal {
		// Jika correctnya sama
		if s[i].Correct == s[j].Correct {

			// Untuk mengecek Nama. Contoh Budi > Agus = A > B (true)
			if s[i].Name > s[j].Name {
				return true
			}
		} else if s[i].Correct < s[j].Correct { // Jika correctnya i lebih kecil dari j
			return true
		}
	} else if aTotal < bTotal { // Jika totalnya i lebih kecil dari j
		return true
	}

	return false
}

func (s Scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Scores) TopStudents() []string {
	// Looping
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1; j++ {
			isLess := s.Less(j, j+1)
			if isLess {
				s.Swap(j, j+1)
			}
		}
	}
	names := []string{}
	for _, score := range s {
		names = append(names, score.Name)
	}
	return names
}

func main() {
	scores := Scores([]Score{
		{"Levi", 3, 2, 2}, //score: 3*4 - 2 = 10
		{"Agus", 3, 4, 0}, //score: 3*4 - 4 = 8
		{"Anon", 3, 0, 4},
	})
	sort.Sort(scores)
	fmt.Println(scores)
	fmt.Println(scores.TopStudents())
}
