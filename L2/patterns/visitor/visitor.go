package main

// Сперва мы определяем интерфейс посетителя следующим способом:

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForTriangle(*Triangle)
}

//Функции visitForSquare(square), visitForCircle(circle), visitForTriangle(triangle)
//позволят нам добавлять функционал для квадратов, кругов и треугольников соответственно.
