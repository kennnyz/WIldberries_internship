package main

import (
	"bufio"
	"fmt"
	"os"
)

//Отсортировать строки в файле по аналогии с консольной утилитой sort
//(man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

//Реализовать поддержку утилитой следующих ключей:

//-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
//-n — сортировать по числовому значению
//-r — сортировать в обратном порядке
//-u — не выводить повторяющиеся строки

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	k := 0
	for scanner.Scan() {
		k++
		scanner.Text()
	}

	var rows [k][1]string
	rows[0][0] = "sdf"

	fmt.Println(k, rows)
}
