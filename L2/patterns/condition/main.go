package main

import (
	"log"
)

//Состояние — это поведенческий паттерн, позволяющий динамически изменять поведение объекта при смене его состояния.
//Поведения, зависящие от состояния, переезжают в отдельные классы.
//Первоначальный класс хранит ссылку на один из таких объектов-состояний и делегирует ему работу.

// condition  design pattern

func main() {
	vendingMachine := NewVendingMachine(10, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	//err = vendingMachine.insertMoney(5)
	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
