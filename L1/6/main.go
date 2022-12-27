//Реализовать все возможные способы остановки выполнения горутины.

// Способ № 1. Горутина останавливается когда канал "quit" становится готовым для работы. Т.е если в quit пришло значение горутина останавливается

package main

import "fmt"

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

// Спопос №2.
