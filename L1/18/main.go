package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	sync.Mutex
}

func (c *Counter) Increment() {
	c.Lock()
	c.count++
	c.Unlock()
}

func main() {
	wg := &sync.WaitGroup{}
	counter := Counter{
		count: 0,
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter.count)
}
