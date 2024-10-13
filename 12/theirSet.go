package main
/* Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество. */

import (
	"fmt"
)

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	stringsSet := make(map[string]struct{})
	for _, word := range strings {
		stringsSet[word] = struct{}{}
	}
	for key, _ := range stringsSet {
		fmt.Println(key)
	} 
}