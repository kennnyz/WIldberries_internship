package main

type Shape interface {
	getType() string
	accept(Visitor)
}

type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

type Circle struct {
	radius int
}

func (s *Circle) accept(v Visitor) {
	v.visitForCircle(s)
}

func (s *Circle) getType() string {
	return "Circle"
}

type Triangle struct {
	height int
	aside  int
}

func (s *Triangle) accept(v Visitor) {
	v.visitForTriangle(s)
}

func (s *Triangle) getType() string {
	return "Triangle"
}
