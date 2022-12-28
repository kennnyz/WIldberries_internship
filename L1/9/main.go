package main

import "fmt"

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
func main() {
	massive := []int{1, 92, 4, 24, 13, 5, 4, 4, 12, 35, 362, 34, 123, 562, 33, 4332, 34, 99, 234, 0, -34, 324, -33}
	//massive := []int{3, 4}
	c1 := make(chan int)
	c2 := make(chan int)

	go reader(massive, c1)
	go redactor(c1, c2)

	for {
		_, open := <-c1
		if !open {
			break
		}
		fmt.Println(<-c2)
	}
}

func reader(ms []int, ch chan int) {
	for _, v := range ms {
		ch <- v
	}
	close(ch)
}

func redactor(reader, ch chan int) {
	for {
		item := <-reader
		ch <- item * 2
	}
}
