package main

import (
	"fmt"
)
/* 
	Реализовать бинарный поиск встроенными методами языка.
*/

// binarySearch выполняет бинарный поиск в отсортированном срезе.
func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	result := -1
	for low <= high {
		mid := low + (high-low)/2 // Вычисляем средний индекс

		if arr[mid] == target {
			result = mid // Элемент найден, возвращаем индекс
		} else if arr[mid] < target {
			low = mid + 1 // Ищем в правой половине
		} else {
			high = mid - 1 // Ищем в левой половине
		}
	}

	return result // Элемент не найден
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7

	result := binarySearch(arr, target)

	if result != -1 {
		fmt.Printf("Элемент %d найден на индексе %d\n", target, result)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
