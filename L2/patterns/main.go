package main

// Использовать лишь один метод
import (
	"fmt"
)

type myType struct {
	Length float64
	Width  float64
}

func (r *myType) Area() float64 {
	return r.Length * r.Width
}

func (r *myType) AnotherFunc() string {
	return fmt.Sprintf("This is rectangle and his area is %v ", r.Area())
}

func (r *myType) SetLength(i float64) {
	r.Length = i
}

type onlyOneMethod struct {
	oko myType
}

func (m onlyOneMethod) AnotherFunc() {
	m.oko.AnotherFunc()
}

func main() {
	pop := onlyOneMethod{}
	pop.AnotherFunc()

}
