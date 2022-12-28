package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := map[int]int{}
	mux := &sync.Mutex{}

	go writeLoop(m, mux)
	go writeLoop(m, mux)
	go writeLoop(m, mux)
	go writeLoop(m, mux)

	go readLoop(m, mux)

	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int, mux *sync.Mutex) {
	for i := 0; i < 10; i++ {
		mux.Lock()
		m[i] = i
		mux.Unlock()
	}
	fmt.Println("Working")
}

func readLoop(m map[int]int, mux *sync.Mutex) {
	mux.Lock()
	for k, v := range m {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(k, "-", v)
	}
	mux.Unlock()
	fmt.Println("Data has been successfully recorded")
	return
}
