package main

import (
	"fmt"
	"strings"
	"unicode"
)

//"a4bc2d5e" => "aaaabccddddde"
//"abcd" => "abcd"
//"45" => "" (некорректная строка)
//"" => ""

func Unpack(s string) string {
	if s == "" {
		return ""
	}
	var result strings.Builder
	var last rune
	var count int
	for _, r := range s {
		if unicode.IsDigit(r) {
			if last == 0 {
				return ""
			}
			count = count*10 + int(r-'0')
		} else {
			if last != 0 {
				if count == 0 {
					count = 1
				}
				result.WriteString(strings.Repeat(string(last), count))
				count = 0
			}
			last = r
		}
	}
	if count == 0 {
		count = 1
	}
	result.WriteString(strings.Repeat(string(last), count))
	return result.String()
}

func main() {
	sad := Unpack("pop2opo6k3t3mnmn6")
	fmt.Println(sad)
}
