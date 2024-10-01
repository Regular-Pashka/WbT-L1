package main

/* 
	Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.



*/

import (
	"fmt"
	"sync"
)

func square(num int, wg *sync.WaitGroup/* , mu *sync.Mutex */) {
	defer wg.Done()
	// defer mu.Unlock()
	result := num * num
	fmt.Println(result)
}

func main() { 
	// nums := make([]int, 0, 2)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	nums := [3]int{1,2,3}
	var wg sync.WaitGroup
	// var mu sync.Mutex
	for _, val := range nums {
		wg.Add(1)
		// mu.Lock()
		go square(val, &wg/* , &mu */);
	}
	wg.Wait()
}
