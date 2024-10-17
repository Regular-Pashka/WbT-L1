package main

import (
	"fmt"
	"strings"
)

/* Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow». */

func reverseWords(s string) string {
	split := strings.Split(s, " ")
	length := len(split)
	for i := 0; i < length / 2; i++ {
		split[i], split[length - i -1] = split[length - i - 1], split[i]
	}
	result := split.
	return result
}

func main() {
	s := "snow dog sun"
	fmt.Println(s)
	fmt.Println(reverseWords(s))
}