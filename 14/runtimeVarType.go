package main
/* 
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

*/

import (
	"fmt"
)
/* 
	в пустом интерфесе может хранится только одно значение в один момент времени
	чтобы узнать
*/

func printType(v interface{}) {
    switch v := v.(type) {
    case int:
        fmt.Println("Это int:", v)
    case string:
        fmt.Println("Это string:", v)
    case bool:
        fmt.Println("Это bool:", v)
    case chan int:
        fmt.Println("Это chan int:", v)
    default:
        fmt.Println("Неизвестный тип")
    }
}

func main() {
	var a interface{}
	a = 2
	printType(a)	
	a = "dsa"
	printType(a)	
	a = false
	printType(a)	
	a = make(chan int)
	printType(a)	
}