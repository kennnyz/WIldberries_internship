package main

import (
	"fmt"
	"strconv"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func main() {
	c := make(chan string)
	go count("sheep", c, 10)
	for msg := range c {
		fmt.Println(msg)
	}

}

func count(thing string, c chan string, n int) {
	for i := 1; i <= n; i++ {
		c <- strconv.Itoa(i) + ":" + thing
		time.Sleep(time.Second * 1)
	}
	close(c)
}
