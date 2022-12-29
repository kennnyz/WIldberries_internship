package main

import (
	"fmt"
	"strings"
	"time"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

// Принцип работы такой же как и в 19 номере
func reverse(s string) string {
	ss := strings.Split(s, " ")
	j := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(ss[i], ss[j])
		ss[i], ss[j] = ss[j], ss[i]
		j--
	}
	ans := strings.Join(ss, " ")
	return ans
}

func main() {
	fmt.Println(reverse("snow dog sun")) // sun dog snow

}
