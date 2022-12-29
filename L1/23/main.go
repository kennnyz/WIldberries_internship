package main

import "fmt"

// Удалить i-ый элемент из слайса.
func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	a := []int{3, 34, 5, 352, 234, 99, 44, 0}
	a = remove(a, 1) // Удалит 34
	fmt.Println(a)
}
