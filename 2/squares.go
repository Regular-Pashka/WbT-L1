package main

/* 
	Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.



*/

import (
	"fmt"
	"sync"
)

func squares(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := num * num
	fmt.Println(result)
}

func main() { 
	nums := make([]int, 0, 2)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	var wg sync.WaitGroup
	// wg.Wait()
	for _, val := range nums {
		wg.Add(1)
		go squares(val, &wg);
	}
	wg.Wait()
}
