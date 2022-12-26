package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов с использованием конкурентных вычислений.

func main() {
	massive := [5]int{2, 4, 6, 8, 10}
	ans := 0

	wg := sync.WaitGroup{}
	for i := range massive {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			ans += x * x
		}(massive[i])
	}
	wg.Wait() // Программа ждет выполнения всех горутин
	fmt.Println(ans)
}
