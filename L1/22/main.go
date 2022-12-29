package main

import (
	"fmt"
	"strings"
	"time"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20 = 1 048 576

func main() {
	var a float64
	var b float64
	var operation string
	for {
		time.Sleep(3 * time.Second)
		fmt.Println("Введите а: ")
		_, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Println("Incorrect input")
			return
		}
		fmt.Println("Введите b: ")
		_, err = fmt.Scanln(&b)
		if err != nil {
			fmt.Println("Incorrect input")
			return
		}
		fmt.Println("Выберите операцию: (* / + -)  ")
		_, err = fmt.Scanln(&operation)
		if err != nil {
			return
		}
		operation = strings.TrimSpace(operation) // Удаление пробелов
		switch operation {
		case "*":
			fmt.Println("Ответ: ", a*b)
		case "/":
			fmt.Println("Ответ: ", a/b)
		case "+":
			fmt.Println("Ответ: ", a+b)
		case "-":
			fmt.Printf("Ответ: %v ", a-b)
		default:
			fmt.Println("error operation")
		}
	}

}
