package main

//Паттерн Посетитель позволяет вам добавлять поведение в структуру без ее изменения.

func main() {
	square := Square{side: 10}
	circle := Circle{4}
	triangle := Triangle{5, 8}

	// our visitor.
	areaCalculator := &AreaCalculator{}
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	triangle.accept(areaCalculator)

	// основной метод который работает без визитора
	println(square.getType())
}
