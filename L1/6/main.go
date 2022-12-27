//Реализовать все возможные способы остановки выполнения горутины.

// Способ № 1. Горутина останавливается когда канал "quit" становится готовым для работы. Т.е если в quit пришло значение - горутина останавливается

package main

import (
	"context"
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
func main() {
	ch := make(chan int, 1)
	ctxTimeout, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the doSomething function
	go doSomething(ctxTimeout, ch, 100)
	time.Sleep(1 * time.Second)

	for {
		select {
		case <-ch:
			fmt.Println("Read from ch ", <-ch)
		case <-time.After(2 * time.Second):
			fmt.Println("Timed out. Took over than 2 second")
			return
		}
	}
}
func fibonaci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonaci(n-1) + fibonaci(n-2)
}

func doSomething(ctx context.Context, ch chan int, n int) {
	for i := 0; i < n; i++ {
		select {
		default:
			ch <- fibonaci(i)
		case <-ctx.Done():
			fmt.Println("Canceled by timeout")
			return
		}
	}
}
