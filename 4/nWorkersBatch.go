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
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range data {
		fmt.Printf("worker %d show us an integer: %d\n", id, val)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	workerCount := 0
	maxBuffer := 100
	i := 0
	data := make(chan int, maxBuffer)
	c := make(chan os.Signal, 1)

	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&workerCount)
	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go worker(w, data, &wg)
	}
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// signalReceived := false
	signalReceived := false
	go func() {
		<-c
		signalReceived = true
		close(data)
	}()
	/*
		Вся проблема по записи в закрытый канал крылась тут. Внутри функции воркер, рэнжом идет пробег по каналу, но в сам канал значения записываются бесконечно, из-за этого рэнж цикл работает бесконечно, соответственно сигнал ctrl+c приходит в канал с, но ждет бесконечность времени, пока воркеры отработают, а они работают по бесконечно записывающему каналу. РЕШЕНИЕ: ДОБАВИТЬ УСЛОВИЕ В ЦИКЛ
	*/
	for !signalReceived {
		time.Sleep(300 * time.Millisecond) //запись в закрытый канал все равно будет если убрать здесь time sleep,
		if signalReceived {                // это условие нужно т.к. в момент получение сигнала тело этого цикла уже может пройти проверку на !signalReceived и быть в отрабатывании такта, с помощью добавления этого if-a мы удостоверимся что в случае чего цикл не будет писать данные в закрытый канал.
			break
		}
		data <- i
		i++
	}
	wg.Wait()
	fmt.Println("Все горутинки отработали")

}

/*
	Вывод:
	Это была задача на воркер-пул


	Вывод 03.10: select во первых выполняет операцию внутри case, это не аналог if, который ждет пока условие будет выполнено в другой части кода. во вторых select из списка case'ов выполняет в первую очередь тот, который является неблокирующей операцией.

*/

/*

	ВЫВОД ОТ 03.10.2024
	Я не понимаю почему программа себя так ведет. Вот Вывод
	====================================================================
	worker 2 finishedworker 3 finishedworker 1 finishedВсе горутинки отработали
	zapis do
	panic: send on closed channel

	goroutine 1 [running]:
	main.main()
			/Users/staciesu/golangStudy/Golang/goKostyaWay/goWbTech/wbTL1/4/nWorkersBatch.go:73 +0x30a
	exit status 2
	====================================================================
	Почему-то после присвоения signalReceived значения true, цикл все равно продолжает писать в канал, и закономерно ловит панику
	Не разобрался: что такое паника
*/


/* 
	ВЫВОД 04.10:
	За эту задачу понял больше о каналах воркерпулах и вэйтгруппах чем за последние три месяца

*/