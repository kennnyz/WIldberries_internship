package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"strconv"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.

//Программа должна завершаться по нажатию Ctrl+C.

func main() {
	N := 3
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		for sig := range signalChannel {
			fmt.Printf("%d workers were killed when captured '%v' signal caused by user\n", N, sig)
			os.Exit(0)
		}
	}()
	jobs := make(chan int, N)

	for i := 0; i < N; i++ {
		go worker(jobs)
	}

	for i := 0; i < 100; i++ {
		jobs <- i
	}
}

func worker(jobs <-chan int) {
	for n := range jobs {
		io.WriteString(os.Stdout, strconv.Itoa(n)+": "+strconv.Itoa(fib(n))+"\n")
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
