package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

*/

func square(in, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// defer mu.Unlock()
	wg.Add(1)
	num := <- in
	result := num * num
	out <- result
	// fmt.Println(result)
}

func sumSquares(out chan int, wg *sync.WaitGroup, sum *int) {
	defer wg.Done()
	wg.Add(1)
	a := <- out
	*sum += a
}

func main() {
	var wg *sync.WaitGroup
	sum := 0
	nums := [5]int{2, 4, 6, 8, 10}
	in := make(chan int)
	out := make(chan int)	
	for _, val := range nums {
		in <- val
		go square(in, out, wg)
		go sumSquares(out, wg, &sum)
	}
	wg.Wait()
	fmt.Println(sum)
}
