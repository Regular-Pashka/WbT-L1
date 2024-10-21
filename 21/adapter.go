package main
/* 
	Реализовать паттерн «адаптер» на любом примере.

	Теория: 
	паттерн адаптер может адаптировать два или более разных, принадлежащих разным структурам по названию, но схожих по сути метода в один, это будет видно на примере с кошкой и собакой. Пример: кошки и собаки по-разному реагируют на одно и тоже событие, кошка - мяукает, собака - гавкает, человек - вскрикивает и т.д. Писать для каждой сущности вызов конкретного метода может быть долго и затратно по времени, подобно тому как затратно четыре раза складывать двойки друг с другом, вместо умножения оных на 4. Адаптер выступает своего рода "умножением", то есть адаптирует некие разные методы так, чтобы их можно было вызывать одной и той же сигнатурой.
	*/

import (
	"fmt"
)

type Dog struct{}

func (dog Dog) Bark() string {
	return "Bark Bark"
}

type Cat struct{}

func (cat Cat) Meow(isSatisfied bool) string {
	if isSatisfied {
		return "Meow Meow"
	}
	return "Angry view"
}

type Peer struct{}

func (human Peer) Reaction() string {
	return "AAAAAAAAAAAAA"
}

// Теперь пишем интерфейс адаптеров. Адаптеры представляют собой структуры содержащие(путем встраивания)

type AliveAdapter interface{
	Reaction() string
}

// После нужно написать адаптер для каждой структуры, которые не удовлетворяют интерфейсу

type DogAdapter struct {
	Dog
}

type CatAdapter struct{
	Cat
}

//Метод-адаптер для каждого адаптируемой структуры

func (adapter CatAdapter) Reaction() string {
	return adapter.Meow(true)
}

func (adapter DogAdapter) Reaction() string {
	return adapter.Bark()
}

// Конструкторы для создания адаптеров

func NewDogAdapter(dog Dog) AliveAdapter {
	return DogAdapter{}
}

func NewCatAdapter(cat Cat) AliveAdapter {
	return CatAdapter{}
}

// Можно ли обойтись без конструкторов? Ответ: можно, просто нужно в мейне объявить переменные структур адаптеров. Зачем тогда нужны конструкторы? Возможно, чтобы не сорить в мейне.....

func main() {
	cat := CatAdapter{}
	dog := DogAdapter{}
	creatures := [3]AliveAdapter{dog, cat, Peer{}}
	// creatures := [3]AliveAdapter{NewDogAdapter(Dog{}), NewCatAdapter(Cat{}), Peer{}}
	for _, alive := range creatures {
		fmt.Println(alive.Reaction())
	}	
}