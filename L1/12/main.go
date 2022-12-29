package main

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

import (
	"fmt"
)

func getUnique(str []string) (res []string) {
	//struct{}{} имеет вес 0 байт, поэтому мы используем ее в качетсве значения в нашей мапе
	m := make(map[string]struct{})
	for _, v := range str {
		m[v] = struct{}{}
	}
	// Так как ключи уникальные в результате цикла получаем массив с уникальными элементами
	for k := range m {
		res = append(res, k)
	}
	return
}

func main() {
	pop := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(getUnique(pop))
}
