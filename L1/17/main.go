package main

import (
	"fmt"
	"sort"
)

// Реализовать бинарный поиск встроенными методами языка
func main() {
	arr := []int{0, 10, 2, 5, 1, 3, 7}
	// Бинарный поиск работает только с отсортированным массивом
	sort.Ints(arr)
	fmt.Println(sort.SearchInts(arr, 10)) // Ответ 6
}
