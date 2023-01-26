package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	res := make(chan interface{})

	defer close(res)
	if len(channels) == 0 {
		return res
	}
	for i := 0; ; i++ {
		if i == len(channels) {
			i = 0
		}
		select {
		case <-channels[i]:
			return res
		default:
		}
	}
}

func main() {
	start := time.Now()
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	<-or(
		sig(22*time.Second),
		sig(10*time.Second),
		sig(5*time.Second),
		sig(12*time.Second),
		sig(13*time.Second),
		sig(11*time.Second),
	)
	fmt.Printf("fone after %v \n ", time.Since(start))
}
