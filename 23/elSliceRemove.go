package main

import "fmt"

/* 
	TODO:
	Удалить i-ый элемент из слайса.
*/

func removeEl(s *[]int, ind int) {
	if ind >= 0 && ind < len(*s) {
		*s = append((*s)[:ind], (*s)[ind + 1:]...)
	}
}

func main() {
	s := make([]int, 10)
	length := len(s)
	var ind int
	fmt.Scanf("%d", &ind)
	for i := 0; i < length; i++ {
		s[i] = i
		fmt.Println(s[i])
	}
	removeEl(&s, ind)
	fmt.Println("=========")
	for _, val := range s{
		fmt.Println(val)
	}
}
