package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var A = flag.Int("A", 0, "print N lines after match")
var B = flag.Int("B", 0, "print N lines before match")
var C = flag.Int("C", 0, "print +- N around match")
var c = flag.Bool("c", false, "print number of matches")
var i = flag.Bool("i", false, "ignore case")
var v = flag.Bool("v", false, "invert search")
var F = flag.Bool("F", false, "exact non-pattern match")
var n = flag.Bool("n", false, "print line index")

func contain(linestring []string, searchString string) bool {
	for _, v := range linestring {
		if searchString == v {
			return true
		}
	}
	return false
}

func grep(pattern string, filename string, A, B, C int, c, i, v, F, n bool) (any, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Будет хранится сама строка где содержится паттерн и номер строки
	ifMeet := make(map[string]int)
	// Количество строк соддержищих паттерн
	counter := 0
	// Слайс из всех строк
	lines := []string{}
	// кол-во строк
	j := 0

	for scanner.Scan() {
		// Получаем каждую отдельную строку
		line := scanner.Text()
		if i {
			line = strings.ToLower(line)
		}
		// Добавляем
		lines = append(lines, line)
		// если у нас в строке встретился паттерн, то добавляем его в список строк, запоминаем мапу и увеличиваем кол-во
		if strings.Contains(line, pattern) {
			ifMeet[line] = j
			counter++
		}
		j++
	}
	if counter == 0 {
		println("There is no pattern")
		return nil, nil
	}

	if A > 0 {
		for _, value := range ifMeet {
			fmt.Println("Line : ", value+1)
			if value <= j-A-1 {
				for k := value; k <= A+value-0; k++ {
					fmt.Println("another step ----> ", lines[k])

				}
				fmt.Println("--------------------")
			} else {
				for l := value; l < j; l++ {

					fmt.Println("another step ----> ", lines[l])
				}
			}
		}
		return nil, nil
	}
	if B > 0 {
		for _, value := range ifMeet {
			if value > B {
				for i := value - B; i <= value-0; i++ {
					fmt.Println("another step -----> ", lines[i])
				}
				fmt.Println("--------------------")
			} else {
				for i := 0; i <= B-1; i++ {
					fmt.Println("another step -----> ", lines[i])
				}
				fmt.Println("--------------------")
			}

		}

		return nil, nil
	}

	if n {
		for i, v := range ifMeet {
			fmt.Println(i, "- ", v+1)
		}

	}
	// флаг с - вывести количество встреч паттерна
	if c {
		fmt.Printf("{%d} this many times {%s} has meeted in this file \n", counter, pattern)
		return nil, nil
	}

	for k := range ifMeet {
		fmt.Println(k)
	}

	return nil, nil
}

func main() {
	flag.Parse()
	_, err := grep("example", "text.txt", *A, *B, *C, *c, *i, *v, *F, *n)
	if err != nil {
		fmt.Println(err)
	}
}
