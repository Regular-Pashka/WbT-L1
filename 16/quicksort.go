package main
/* Реализовать быструю сортировку массива (quicksort) встроенными методами языка. */

import (
	"fmt"
)

func quicksort(arr []int, left int, right int) {
	if len(arr) < 2 {
		return 
	}
	pivot := arr[right]
	j := (left - 1)

	for i := left; i < right; i++ {
		if arr[i] < pivot {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
	}
	arr[j + 1], arr[right] = arr[right], arr[j + 1]

	quicksort(arr, left, j - 1)
	quicksort(arr, j + 1, right)
}

func main() {
	someSlice := []int{1, 5, 2, 3, 64, 213, 543, 21, 423, 65, -11, 432, 0}
	fmt.Println(someSlice)
	quicksort(someSlice, 0, len(someSlice) - 1)
	fmt.Println(someSlice)
}