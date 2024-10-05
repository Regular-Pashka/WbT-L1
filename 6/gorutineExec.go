package main
/* 

	6. Реализовать все возможные способы остановки выполнения горутины. 

	ВОПРОС: Какие есть способы остановки выполнения горутины?
	1. Через передачу сигнального канала

*/
import (
	"fmt"
	"sync"
)


func multByTwo(id int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range in {
		out <- val * 2
		fmt.Printf("worker %d doubled number %d", id, val)
	}
	fmt.Printf("worker %d finished his work", id)
}

func main() {
	var wg sync.WaitGroup
	in := make(chan int, 30)
	out := make(chan int, 30)
	workerCount := 3
	
	for i := 0; i < 100; i++ {
		in <- i
	} //заполнение входного канала
	for w := 0; w < workerCount; w++ {
		wg.Add(1)
		go multByTwo(w, in, out, &wg)
	}
	for doubled := range out {
		fmt.Print(doubled)
	}
	wg.Wait()
	close(in)
	close(out)
	fmt.Print("DONE")
}