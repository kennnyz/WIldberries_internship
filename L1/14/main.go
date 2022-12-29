package main

import (
	"fmt"
	"reflect"
)

// Разработать программу,
// которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
func main() {
	var str string
	var integer int
	var chan1 interface{} = make(chan int)
	s := true

	//с помощью switch кострукции которая продемонстрирована в фукции tip()
	fmt.Println(tip(str), tip(integer), tip(chan1), tip(s))
}

func tip(a any) string {
	switch a.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case chan int:
		return "chan"
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println(reflect.TypeOf(a).String())
	}

	return "<nil>"
}
