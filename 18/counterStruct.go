package main

/*
	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
*/

import (
	"fmt"
	"sync"
)

type Counter struct {
	counter int
	mu sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	c := Counter{counter:0}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.mu.Lock()
			for j := 0; j < 100; j++ {
				c.counter++
			}
			c.mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(c.counter)
}