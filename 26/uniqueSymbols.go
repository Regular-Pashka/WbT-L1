package main

import (
	"fmt"
	"strings"
)

/*
	TODO:
	Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:

	abcd — true

	abCdefAaf — false
	aabcd — false
*/

func isUnique(s string) bool{
	uniqueFlag := true
	s = strings.ToLower(s)
	uniqueCheck := make(map[rune]int)
	for _, symbol := range s {
		uniqueCheck[symbol]++
	}
	for _, symbAmount := range uniqueCheck {
		if symbAmount > 1 {
			uniqueFlag = false
			break
		}
	}
	return uniqueFlag
}

func main() {
	a := "privet2"
	fmt.Println(isUnique(a))
}