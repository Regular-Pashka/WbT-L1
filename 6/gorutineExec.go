package main

/*

	6. Реализовать все возможные способы остановки выполнения горутины.

	ВОПРОС: Какие есть способы остановки выполнения горутины?
	1. Через передачу сигнального канала
	2. Через контекст
		2.1. withCancel
		2.2. withTimeout
		2.3. withDeadline
		Все это одно и тоже

*/
import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
1. Способ
func square(num int, signalChan chan bool, wg *sync.WaitGroup) {

defer wg.Done()
select {
	case <-signalChan:
		fmt.Printf("Goroutine stopped because of signal chan, num: %d\n",  num)
		return
	default:
	}
	result := num * num
	fmt.Println(result)
}
*/

//2.1. With Cancel
/* func sumOfChan(ctx context.Context, in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("gourutine was stopped by context")
			return
		case num, ok :=<- in:
			if !ok {
				fmt.Println("gourutine finished her work", sum)
				return
			}
			sum += num
		}
	}
}
*/

func sumOfChan(ctx context.Context, in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("gourutine was stopped by context")
			return
		case num, ok :=<- in:
			if !ok {
				fmt.Println("gourutine finished her work", sum)
				return
			}
			sum += num
		}
	}
}

func main() {
	//2.2 With Timeout
	var wg sync.WaitGroup
	c := make(chan int, 20)
	ctx, cancel := context.WithTimeout(context.Background(), 500 * time.Millisecond)
	wg.Add(1)
	go sumOfChan(ctx, c, &wg)
	for i := 0; i < 100; i++ {
		c <- i
	}
	cancel()
	close(c)
	wg.Wait()
	




	// 2.1 WithCancel
/* 
	var wg sync.WaitGroup
	c := make(chan int, 20)
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go sumOfChan(ctx, c, &wg)
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
	cancel()
	wg.Wait() */


	
	// 1. Использование закрывающего канала 
	/*
	
	var wg sync.WaitGroup
	num := 5
	signalChan := make(chan bool)
	wg.Add(1)
	go square(num, signalChan, &mu , &wg);
	signalChan <- false
	wg.Wait()
	 
	*/
}