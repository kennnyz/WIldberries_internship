package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func Intersection(arr1, arr2 []int) []int {
	ans := []int{}
	// Запись только уникальных значений в 1 массиве
	map1 := make(map[int]int)
	for _, v := range arr1 {
		map1[v] = 1
	}

	for _, v := range arr2 {
		if map1[v] == 1 {
			map1[v] += 1
			ans = append(ans, v)
		}
	}
	return ans
}

func main() {
	mn1 := []int{5, 5, 1, 23}
	mn2 := []int{4, 23, 5, 5, 2}

	fmt.Println(Intersection(mn1, mn2))
	// пересечение - (5,23)

}
