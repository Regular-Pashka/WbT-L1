package main

/* 
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

import (
	"fmt"
)

func reverseString(s string) string{
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length / 2; i++ {
		runes[i], runes[length - i - 1] = runes[length - i - 1], runes[i]
	}
	return string(runes)
}

func main() {
	s := "главрыба"
	fmt.Println(s)
	fmt.Println(reverseString(s))
}