package main

// Посетитель

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForTriangle(*Triangle)
}
