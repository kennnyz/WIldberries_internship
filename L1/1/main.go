package main

import "fmt"

//Human  - родительсякая структура

type Human struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

// Конструктор для создания Human

func NewHuman(name string, age int, sex string, weight, height int) *Human {
	return &Human{name, age, sex, weight, height}
}

// Метод который реализует Human
func (h *Human) isAdult() bool {
	return h.age >= 18
}

// Action дочерняя структура Human которая наследует его (Human) поля и методы
// Прикол Action в том что - эта структура во многом повторяет Human. Этой структуре придется описывать только доп. функциональность

type Action struct {
	Human        // Встроенный тип
	job   string // Регулярное поля
	alone bool   //Регулярное поле
}

// Функция NewAction — это конструктор, который использует свои
//параметры для создания Action со встроенным значением Human

func NewAction(name string, age int, sex string, weight int, height int, job string, alone bool) *Action {
	return &Action{
		*NewHuman(name, age, sex, weight, height), job, alone,
	}
}

func (a *Action) isFat() bool {
	if !a.isAdult() {
		return false
	}
	return a.weight > 90
}

func main() {
	person := NewAction("m", 1, "male", 92, 192, "OZON", true)
	fmt.Println(person.isFat())

	h := Human{name: "aa"}
	a := Action{
		Human: h,
		job:   "aa",
	}

	b := Action{
		Human: h,
		job:   "aa",
	}
	a.name = "bbb"
	fmt.Println(b.name)
}
