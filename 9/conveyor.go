package main

/*
	TODO:
	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	nums := []int{1, 2, 3, 4, 5, 6, 2, 32, 1}
	numsChan := make(chan int)
	doubledNumsChan := make(chan int)
	wg.Add(3)
	go func() {
		defer wg.Done()
		for _, num := range nums {
			numsChan <- num
		}
		close(numsChan)
	}()
	go func() {
		defer wg.Done()
		for num := range numsChan {
			doubledNumsChan <- num * 2
		}
		close(doubledNumsChan)
	}()
	go func() {
		defer wg.Done()
		for doubledNum := range doubledNumsChan {
			fmt.Println(doubledNum)
		}
	}()
	wg.Wait()
}