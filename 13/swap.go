package main
/* Поменять местами два числа без создания временной переменной. */

import (
	"fmt"
)

func main() {
	a := 2
	b := 3
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}