package main

/*
	5. Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

import (
	"fmt"
	"sync"
	"time"
)


func writeChan(c chan<- int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Print("Запись в канал завершена!")
	i := 0
	for {
		select {
		case <-done:
			close(c)
			return
		case c <- i:
			time.Sleep(time.Millisecond * 100)
			i++
		}
	}
}

func readChan(c <-chan int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Print("Чтение из канала завершено!\n")
	for {
		select {
		case <-done:
			fmt.Println("Time has come! Bye.") 
			return
		case c := <-c:
			fmt.Println(c)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	c := make(chan int, 50)
	doneSignal := make(chan struct{})
	seconds := 0
	fmt.Print("Введите количество секунд, по истечению которых программа прекратит работу: ")
	fmt.Scan(&seconds)
	wg.Add(2)
	go func() {
		time.Sleep(time.Duration(seconds) * time.Second)
		close(doneSignal)
	}()
	go writeChan(c, doneSignal, &wg)
	go readChan(c, doneSignal, &wg)
	wg.Wait()
}

/* 
	 Вывод по заданию:
	1. Каналы можно использовать как флаги какого-то события, чтобы через select прекращать выполнение каких либо функций
	2. При чтении из закрытого пустого канала возвращается нулевое значение соответствующего типа
	3. defer засовываются в стек и выполняются по принципу последний зашедший - первый выполняемый
	4. вместо doneSignal можно было спользовать time.Timer(), он возвращает структуру Timer в которой есть поле(канал на чтение) T, который бы делал тоже самое
*/