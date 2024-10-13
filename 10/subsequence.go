package main

import (
	"fmt"
	"sort"
)

/* TODO:
	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

	Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

type group map[int][]float64

func main() {
	var intPart int
	var keys []int

	ex := make(group)	
	nums := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Printf("")
	for _, num := range nums {
		intPart = int(num)
		groupKey := (intPart / 10) * 10
		ex[groupKey] = append(ex[groupKey], num)
	}

	for key, _ := range ex {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Printf("%d:%v,\n", key, ex[key])
	}
}
