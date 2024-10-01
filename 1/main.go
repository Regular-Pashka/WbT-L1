package main

import (
	"fmt"
)
/* 
	Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	Height int
	Name string
	Age int
}

type Action struct {
	Human
	Name string
}

func (chel *Human) ChangeName(newName string) {
	chel.Name = newName
} 

func main() {
	act := Action{
		Human: Human{
			Height: 11,
			Name: "SSS",
			Age: 23,
		},
		Name: "AAA",
	}
	act.ChangeName("Паша")
	fmt.Println(act.Name)
	fmt.Println(act.Human.Name)
}

/* Результат:
	ААА
	Паша
	Выводы по заданию: методы можно встраивать от одной структуры к другой. При встраивании методов метод меняет(если меняет) поля своей родительской структуры. 	act.ChangeName("Паша") и 	act.Human.ChangeName("Паша") это одно и тоже.


*/