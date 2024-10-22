package main

import (
	"fmt"
	"time"
)

/*
	TODO:
	Реализовать собственную функцию sleep.
*/

func Sleep(seconds int) {
	timer := time.NewTimer(time.Second * time.Duration(seconds))
	<-timer.C
}

func main() {
	fmt.Println(3)
	Sleep(5)
	fmt.Println(4)
}