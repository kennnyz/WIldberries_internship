package shape

import "fmt"

type Shape interface {
	Area() float64
	AnotherFunc() string
	SetLength(i float64)
}

type Rectangle struct {
	Length float64
	Width  float64
}

func NewRectangle(length float64, width float64) Shape {
	return &Rectangle{Length: length, Width: width}
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r *Rectangle) AnotherFunc() string {
	return fmt.Sprintf("This is rectangle and his area is %v ", r.Area())
}

func (r *Rectangle) SetLength(i float64) {
	r.Length = i
}
