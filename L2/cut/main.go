package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flagf := flag.Int("f", 0, " - выбрать поля")
	flagd := flag.String("d", "\t", "использовать другой разделитель")
	flags := flag.Bool("s", false, "только строки с разделителем ")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')

		delimLine := []string{}
		if strings.Contains(input, *flagd) {
			delimLine = strings.Split(input, *flagd)
		}

		if *flagf > 0 {
			fmt.Printf("You working with f flag and here is your %d field ---> %s", *flagf, delimLine[*flagf-1])
		}

		if *flags {
			fmt.Println("You working with s flag and here is your string ---> ", input)
		}
	}

}

//
