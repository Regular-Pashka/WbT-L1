package main

import (
	"fmt"
)

/* 
	8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func setBit(num *int64, ind int, bit int) {
	if ind >= 0 && ind < 64 {
		mask := int64(0b1 << ind)
		if bit == 0 {
			*num &= *num & (^mask)
		} else if bit == 1{
			*num |= mask
		}
	}
}

func printBits(num int64) {
	for i := 63; i >= 0; i-- {
		// Проверяем, установлен ли i-й бит
		if (num & (1 << i)) != 0 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
	fmt.Println()
}

func main() {
	num := int64(5)
	fmt.Println(num)
	printBits(num)
	setBit(&num, 3, 1)
	fmt.Println(num)
	printBits(num)
}