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
}

//---------------------------------------------------------------------------------------------------------------------//

// Спопоб №2. Можно остановить горутину с помощью Context
// Если время вычисление числа фибоначи займет больше чем N секунд до горутина остановится

func fibonaci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonaci(n-1) + fibonaci(n-2)
}
func operation1(ctx context.Context) error {
	// Let's assume that this operation failed for some reason
	// We use time.Sleep to simulate a resource intensive operation
	time.Sleep(2 * time.Second)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	for i := 0; i < 100; i++ {
		select {
		default:
			fmt.Println(i, ": ", fibonaci(i))
		case <-ctx.Done():
			fmt.Println("Took time over than 2 second")
			return
		}
	}

}

func main() {
	// Create a new context
	ctx := context.Background()
	// Create a new context, with its cancellation function
	// from the original context
	ctx, cancel := context.WithCancel(ctx)
	// Run two operations: one in a different go routine
	go func() {
		err := operation1(ctx)
		// Если это операция вернет ошибку
		// то закроются все операцию использующие данный контекст
		if err != nil {
			cancel()
		}
	}()
	// Run operation2 with the same context we use for operation1
	operation2(ctx)
}
