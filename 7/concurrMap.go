package main

import (
	// "fmt"
	// "hash/adler32"
	// "strconv"
	"fmt"
	"sync"
)

/*
	Реализовать конкурентную запись данных в map.
*/


func main() {
	/* 
		Можно сделать через sync.Map, но камоон
	*/
	var mu sync.Mutex
	var wg sync.WaitGroup
	mapa := make(map[int]int)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			fmt.Printf("%d gourutine: ", i)
			mapa[i] = i
			fmt.Println(mapa[i] + '\n')
			mu.Unlock()
		}(i)
	}
	wg.Wait()
}