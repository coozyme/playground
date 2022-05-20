// Mengecek jika dua string adalah anagram
// Anagram adalah kata yang dibentuk melalui penataan ulang huruf-huruf dari beberapa kata lain.
//
// Contoh 1
// Input: a = "keen", b = "knee"
// Output: "Anagram"
// Explanation: Jika ditata, "knee" dan "keen" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 2
// Input: a = "fried", b = "fired"
// Output: "Anagram"
// Explanation: Jika ditata, "fried" dan "fired" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 3
// Input: a = "apple", b = "paddle"
// Output: "Bukan Anagram"
// Explanation: Jika ditata, "apple" dan "paddle" memiliki huruf-huruf yang berbeda

package main

import "fmt"

func main() {
	var str1 = "fried"
	var str2 = "fired"
	fmt.Println(AnagramsChecker(str1, str2))
}

func AnagramsChecker(str1 string, str2 string) string {
	// TODO: replace this
	strMap := make(map[string]string)
	var result string

	if len(str1) != len(str2) {
		result = "Bukan Anagram"
	}
	for i := 0; i < len(str1); i++ {
		strMap[str1[:1]] = str1[:1]
	}

	for i := 0; i < len(str2); i++ {
		if str2[:1] == strMap[str2[:1]] {
			result = "Anagram"
		} else {
			result = "Bukan Anagram"
		}
	}

	return result
}
