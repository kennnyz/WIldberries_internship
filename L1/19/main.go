package main

import (
	"fmt"
	"strings"
	"time"
)

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.

// Реализовал функцию, которая работает O(n/2). начало и конец заменяются друг другом и круг сужается
func reverse(s string) string {
	ss := strings.Split(s, "")
	j := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(ss[i], ss[j])
		ss[i], ss[j] = ss[j], ss[i]
		j--
	}
	ans := strings.Join(ss, "")
	return ans
}

func main() {
	fmt.Println(reverse("главрыба")) // абырвалг
}
