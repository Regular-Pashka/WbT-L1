package main
/* 
	4. Реализовать постоянную запись данных в канал (главный поток). 
	Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. 
	Необходима возможность выбора количества воркеров при старте.

	====Не сделано===
	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
	====Не сделано====

	НЕ ПОНИМАЮ ПОЧЕМУ ПРОГРАММА НЕ РЕАГИРУЕТ НА CTRL+C

*/



import (
	"fmt"
	"sync"
	// "bufio"
	"os"
	"syscall"
	"os/signal"
	// "math/rand"
	"time"
)

func worker(id int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range data {
		fmt.Printf("worker %d show us a random: %d\n", id, val)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	workerCount := 0
	maxBuffer := 100
	i := 0
	data := make(chan int, maxBuffer)
    fmt.Print("Введите количество воркеров: ")
    fmt.Scan(&workerCount)
	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go worker(w, data, &wg)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	signalReceived := false
	// go func() {
	// 	<-c
	// 	close(data)
	// }()
	for !signalReceived {
		select {
		case <-c:
			close(data)
			signalReceived = true	
		default:
			i++
			data <- i
		}
	}

	wg.Wait()
	fmt.Println("Все горутинки отработали!")
}


/* 
	Вывод:
	Это была задача на воркер-пул

*/