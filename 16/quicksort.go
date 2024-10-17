package main
/* Реализовать быструю сортировку массива (quicksort) встроенными методами языка. */

import (
	"fmt"
)

// quickSort выполняет быструю сортировку на срезе целых чисел.
func quickSort(arr []int, low, high int) {
	if low < high {
		// Разделяем массив и получаем индекс опорного элемента
		pi := partition(arr, low, high)

		// Рекурсивно сортируем элементы до и после опорного элемента
		quickSort(arr, low, pi - 1)
		quickSort(arr, pi + 1, high)
	}
}

// partition разделяет массив и возвращает индекс опорного элемента
func partition(arr []int, low, high int) int {
	pivot := arr[high] // выбираем последний элемент в качестве опорного
	i := low - 1       // индекс меньшего элемента

	for j := low; j < high; j++ {
		// Если текущий элемент меньше или равен опорному
		if arr[j] <= pivot {
			i++ // увеличиваем индекс меньшего элемента
			arr[i], arr[j] = arr[j], arr[i] // меняем местами
		}
	}
	// меняем местами опорный элемент с элементом, следующим за меньшим
	arr[i + 1], arr[high] = arr[high], arr[i + 1]
	return i + 1
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	n := len(arr)

	quickSort(arr, 0, n - 1)

	fmt.Println("Отсортированный массив:", arr)
}