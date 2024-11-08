package main

import (
	"fmt"
	// "unsafe"
	"sync"
)
/* 
	12. Что выведет данная программа и почему?
*/

// func main() {
// 	n := 0
// 	if true {
// 		n := 1
// 		n++
// 		fmt.Println(n)
// 	}
// 	fmt.Println(n)
// }


/* 
	ВЫВОД ПО ЗАДАЧЕ:
	В языке Go любые блоки кода, ограниченные фигурными скобками {}, создают свою локальную область видимости.
*/ 


/* 13. Что выведет данная программа и почему? */

// func someAction(v []int8, b int8) {
// 	v[0] = 100
// 	v = append(v, b)
// 	fmt.Printf("%p\n", v)
// 	fmt.Println(v)
// }
	
// func main() {
// 	var a = make([]int8, 5, 10)
// 	for i := 0; i < 5; i++ {
// 		a[i] = int8(i + 1)
// 	}
// 	someAction(a, 6)
// 	fmt.Printf("%p\n", a)
// 	p := &a[0]	
// 	for i := 0; i < len(a); i++ {
//         // Вычисляем адрес i-го элемента
//         addr := uintptr(unsafe.Pointer(p)) + uintptr(i)*unsafe.Sizeof(a[0])
//         value := *(*int8)(unsafe.Pointer(addr)) // Преобразуем обратно в указатель на int и разыменовываем
//         fmt.Printf("Value at index %d: %d\n", i, value)
//     }
// }

/* 
	ВЫВОД :
	append в этом случае вернет структуру с тремя полями: тем же указателем, тем же капасити но другим length равным шести. Он действительно добавит по мейновскому указателю шестерку, но так как в мейновском слайсе у нас структура с длиной 5, а не 6, то бишь отличный экземпляр структуры от того, что вернул append в функции, мы эту шестерку можем увидеть только с помощью адресной арифметики(пакета unsafe)
*/

//

// func main() {
// 	slice := []string{"a", "a"}
// 	func(slice []string) {
// 		slice = append(slice, "a") // тут из-за того что в слайсе недостаточно капасити будет создан новый базовый массив. Так как он создан внутри анонимной функциии, slice будет отличаться от мейновского слайса. Тут же, в это анонимной функции будет замена первых двух букв а нового слайса на b. Итого вывод будет [b, b, a][a, a]
// 		slice[0] = "b"
// 		slice[1] = "b"
// 		fmt.Print(slice)
// 	}(slice)
// 	fmt.Print(slice)
// }

// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(wg sync.WaitGroup, i int) { // тут просто вейтгруппу передавали не по указателю из-за чего горутины 
// 			fmt.Println(i)
// 			wg.Done()
// 		}(wg, i)
// 	}
// 	wg.Wait()
// 	fmt.Println("exit")
// }



/* Ниже интересная задачка от тимлида */

// // You can edit this code!
// // Click here and start typing.
// package main

// import "fmt"

// // Вопрос: Объяснить причину ошибки
// func main() {
// 	var counter Counter

// 	counter = NewCounter()
// 	// проверка счетчика
// 	if counter == nil {
// 		fmt.Println("counter is nil")
// 		return
// 	}

// 	fmt.Printf("start count with counter[%v]\n", counter)

// 	for i := 0; i < 5; i++ {
// 		fmt.Println(counter.Count())
// 	}
// }

// type Counter interface {
// 	Count() int
// }

// type SomeCounter struct {
// 	cnt int
// }

// func (c *SomeCounter) Count() int {
// 	c.cnt++
// 	return c.cnt
// }

// func NewCounter() *SomeCounter {
// 	return nil
// }