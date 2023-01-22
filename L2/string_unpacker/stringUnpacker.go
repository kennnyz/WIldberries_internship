package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//"a4bc2d5e" => "aaaabccddddde"
//"abcd" => "abcd"
//"45" => "" (некорректная строка)
//"" => ""

func UnpackingString(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	str := []rune(s)

	if unicode.IsDigit(str[0]) {
		return "", errors.New("invalid string")
	}

	var res strings.Builder

	for i := range str {
		if unicode.IsDigit(str[i]) {
			//ищу правую границу числа
			right := i + 1
			for right < len(str) && unicode.IsDigit(str[right]) {
				right++
			}
			count, err := strconv.Atoi(string(str[i:right]))
			if err != nil {
				return "", err
			}
			for ; count > 1; count-- {
				res.WriteString(string(str[i-1]))
			}
			i = right - 1
		} else {
			if str[i] == '\\' {
				i++
			}
			res.WriteString(string(str[i]))
		}
	}
	return res.String(), nil
}

func main() {
	sad, _ := UnpackingString("qw//5")
	fmt.Println(sad)
}
