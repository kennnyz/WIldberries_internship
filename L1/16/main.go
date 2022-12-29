package main

import (
	"fmt"
	"sort"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
func main() {
	massive := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	sort.Ints(massive)
	fmt.Println(massive)

	family := []struct {
		Name string
		Age  int
	}{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}
	sort.Slice(family, func(i, j int) bool {
		return family[i].Age > family[j].Age
	})

	fmt.Println(family)
}
