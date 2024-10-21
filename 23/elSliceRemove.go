package main

import "fmt"

/* 
	TODO:
	Удалить i-ый элемент из слайса.
*/

func main() {
	s := make([]int, 10)
	length := len(s)
	for i := 0; i < length; i++ {
		s[i] = i
	}
	for _, val := range s {
		fmt.Println(val)
	}
}