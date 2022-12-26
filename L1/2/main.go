package main

import (
	"io"
	"os"
	"strconv"
	"sync"
)

//Написать программу, которая конкурентно
//рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

// Конкурентно имеется ввиду - использовать горутины
// Горутина (goroutine) — это функция, выполняющаяся конкурентно с другими горутинами в том же адресном пространстве.

func main() {
	massive := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for i := range massive {
		wg.Add(1)
		go func(x int) {
			defer wg.Done() // Обозночается что горутина выполнена
			// Простая запись строки
			io.WriteString(os.Stdout, strconv.Itoa(x*x)+"\n")
		}(massive[i])
	}
	wg.Wait() // Программа ждет выполнения всех горутин
}
