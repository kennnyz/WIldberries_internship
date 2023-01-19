package main

import "fmt"

func main() {
	normalBuilder := getBuilder("igloo")

	director := NewDirector(normalBuilder)

	house := director.buildHouse()

	fmt.Printf("Normal house door [%s]\nNormal house floor[%d]\nNormal house window[%s]", house.doorType, house.floor, house.windowType)
}
