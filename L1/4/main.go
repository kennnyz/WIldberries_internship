package main

import (
	"io"
	"os"
	"strconv"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.

//Программа должна завершаться по нажатию Ctrl+C.

func main() {
	N := 2

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 0; i < N; i++ {
		go worker(jobs, results)
	}

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 100; i++ {
		io.WriteString(os.Stdout, strconv.Itoa(i)+": "+strconv.Itoa(<-results)+"\n")
	}
}

// Только получаем с jobs chanel и только отправляем в results chanel
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
