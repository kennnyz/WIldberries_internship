package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fields := flag.Int("f", 0, " - выбрать поля")
	delim := flag.String("d", "\t", "использовать другой разделитель")
	flags := flag.Bool("s", false, "только строки с разделителем ")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		if *flags && !strings.Contains(input, *delim) {
			continue
		}
		delimLine := strings.Split(input, *delim)
		if *fields > 0 {
			fmt.Printf("You working with f flag and here is your %d field ---> %s", *fields, delimLine[*fields-1])
		} else if *fields > len(delimLine) {
			fmt.Println("Field number is more than number of columns")
		}
	}
}

//
