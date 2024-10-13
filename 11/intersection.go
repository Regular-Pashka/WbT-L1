package main
/* Реализовать пересечение двух неупорядоченных множеств. */

import (
	"fmt"
)

func main() {
	// нужно вывести массив в котором есть только входящие в оба массива элементы. Как найти входят ли оба массива в элементы? Нужна функция поиска значения в массиве, то есть я бегу по наибольшему(?) массиву, встречаюсь в элемент 
	nums1 := []int{1, 123, 7456, 32, 12, 43, 11, 64, 23}
	nums2 := []int{1, 33, 122, 7456, 32, 112, 44, 21, 11, 43, 23, 43}
	uniqueSet := make(map[int]struct{})
	result := []int{}

	for _, num := range nums1 {
		uniqueSet[num] = struct{}{}
	}
	for _, num := range nums2 {
		if _, exists := uniqueSet[num]; exists {
			result = append(result, num) 
		}
	}
	for _, num := range result {
		fmt.Println(num)
	}
}