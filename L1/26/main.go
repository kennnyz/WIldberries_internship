package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные
//(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

func isUniqueString(s string) bool {
	ss := strings.Split(s, "")
	m := make(map[string]int)
	for _, v := range ss {
		if m[v] > 1 {
			return false
		}
		m[v] = 1
		m[v]++
	}
	return true
}

func main() {
	fmt.Println(isUniqueString("abB"))       //true
	fmt.Println(isUniqueString("abb"))       //false
	fmt.Println(isUniqueString("abCdefAaf")) //false
	fmt.Println(isUniqueString("jaosfb"))    //true
}
