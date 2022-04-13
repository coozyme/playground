package main

import (
	"fmt"
	"strings"
)

// Salah satu problem populer yang dapat diselesaikan dengan menggunakan Stack adalah mengecek validitas tanda kurung.
// Diberikan sebuah string yang hanya terdapat karakter '(', ')', '{', '}', '[', dan ']'.
// Tentukan apakah sebuah string merupakan sekumpulan tanda kurung yang valid.
// String tanda kurung yang valid adalah
// 1. Tanda kurung buka harus ditutup oleh pasangannya.
// 2. Tanda kurung buka harus ditutup di urutan yang benar.

// Contoh 1
// Input: s = "()"
// Output: true
// Penjelasan: tanda kurung buka '(' ditutup dengan pasangannya yaitu ')'.

// Contoh 2
// Input: s = "{)"
// Output: false
// Penjelasan: tanda kurung buka '{' tidak ditutup oleh pasangannya.

// Contoh 3
// Input: s = "([])"
// Output: true
// Penjelasan: tanda kurung buka ditutup dengan pasangannya dan sesuai dengan urutan.

func IsValidParentheses(s string) bool {
	// TODO: answer here
	data := strings.Split(s, "")
	if len(data) == 0 || len(data)%2 == 1 {
		return false
	}

	for idx, val := range data {
		if (idx+1)%2 == 0 {
			fmt.Println(data[idx-1], val)
			if data[idx-1] == "{" && val == "}" {
				continue
			} else if data[idx-1] == "(" && val == ")" {
				continue
			} else if data[idx-1] == "[" && val == "]" {
				continue
			}

			return false

		}
	}

	return true
}
