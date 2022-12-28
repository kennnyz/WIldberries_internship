package main

import "fmt"

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов.
//Последовательность в подмножноствах не важна.

func main() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 33, 992, -12, 3, 0, 2, 213, 214}
	ans := make(map[int][]float32)
	for _, v := range temperatures {
		step := int(v) - int(v)%10
		ans[step] = append(ans[step], v)
	}
	fmt.Println(ans)
}
