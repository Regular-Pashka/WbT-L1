package main

/*
	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

	var justString string
	func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	}

	func main() {
	someFunc()
	}
*/

import (
	"fmt"
	"strings"
)

// "fmt"

var justString string

func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() {
	v := strings.Repeat("РОССИЯ", 100)//тут задается строка размером 1024 байта
	fmt.Println(v)
	justString = v[500:] // тут происходит копирование данных из среза байтов v в justString
	fmt.Printf("\n\n%s", justString)
}

func main() {
	someFunc()
}