package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func sleeper(i time.Duration) {
	t := time.Now()
	for {
		if time.Now().After(t.Add(i)) {
			return
		}
	}
}
func main() {
	sleeper(5 * time.Second)
	fmt.Println("Sent after 5 sec")
}
