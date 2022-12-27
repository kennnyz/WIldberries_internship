//Реализовать все возможные способы остановки выполнения горутины.

// Способ № 1. Горутина останавливается когда канал "quit" становится готовым для работы. Т.е если в quit пришло значение - горутина останавливается

package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func fib(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("QUIT")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(<-c)
		} // Посчитает первые 100 чисел фибоначи, после чего канал quit получит значение 0 и тем самым горутина остновится
		quit <- 0
	}()
	fib(c, quit)

//---------------------------------------------------------------------------------------------------------------------//

// Спопоб №2. Можно остановить горутину с помощью Context
func longRunningProcess(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return errors.New("failed: Process took too long")
}
func fibonaci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonaci(n-1) + fibonaci(n-2)
}
func worker(ctx context.Context, n int) {
	fibonaci(n)
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("")
	case <-ctx.Done():
		fmt.Println("Took too long")
	}

}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		err := longRunningProcess(ctx)
		if err != nil {
			cancel()
		}
	}()
	worker(ctx, 100)
}
