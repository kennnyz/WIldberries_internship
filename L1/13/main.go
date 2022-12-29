package main

import "fmt"

// Поменять местами два числа без создания временной переменной

// Способ 1. Сложение и вычитание
func swipe1(a, b int) {
	fmt.Println("Was: ", a, b)
	a += b
	b = a - b
	a -= b
	fmt.Println("Become: ", a, b)
}

// Способ 2. Умножение и деление
func swipe2(a, b int) {
	fmt.Println("Was: ", a, b)
	a = a * b
	b = a / b
	a = a / b
	fmt.Println("Become: ", a, b)
}

// Способ 3. Использование побитового исключающего ИЛИ
// XOR манипулирует битами. Возвращает 1, когда значения переменных различны, и 0 в противном случае.
func swipe3(a, b int) {
	fmt.Println("Was: ", a, b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("Become: ", a, b)

	// На примере чисел 10 и 1
	//4-ех битное представление 10 -> 1010
	//
	//4-ех битное представление 1 -> 0001
	//Тогда:
	//На 4 строке: a = a ^ b => 1010 ^ 0001 => 1011 => 7
	//На 5 строке: b = a ^ b => 1011 ^ 0001 => 1010 => 10
	//На 6 строке: a = a ^ b => 1011 ^ 1010 => 0001 => 1
}

// И пожалуй самый простой способ поменять два числа без доп переменной.
func swipe4(a, b int) {
	fmt.Println("Was: ", a, b)
	a, b = b, a
	fmt.Println("Become: ", a, b)
}
func main() {
	var a int = 5
	var b int = 7

	swipe4(a, b)
}
