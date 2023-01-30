package main

import "fmt"

// 1. Определяем интерфейс Strategy  с методов doSearch. Данный метод определяет общее поведение для конкретных алгоритмов, реализующих разные стратегии.

type Strategy interface {
	doSearch(filters map[string]int)
}

type firstAlgorithm struct {
}

func (f *firstAlgorithm) doSearch(filters map[string]int) {
	fmt.Println("First implements strategy ", filters)
}

type secondAlgorithm struct {
}

func (f *secondAlgorithm) doSearch(filters map[string]int) {
	fmt.Println("Second implements strategy ", filters)
}

// напишем класс контекст

type privateValues struct {
	// текущая активная конкретная стратегия
	actionStrategy Strategy
	// текущие фильтры клиента
	userFilters map[string]int
}

func initStrategy(s Strategy) *privateValues {
	//init user filters
	f := make(map[string]int)
	return &privateValues{
		actionStrategy: s,
		userFilters:    f,
	}
}

// установка новой стратегии
func (p *privateValues) setStrategy(s Strategy) {
	p.actionStrategy = s
}

// получение данных по фильтрам
func getData(v *privateValues) {
	v.actionStrategy.doSearch(v.userFilters)
}

func main() {
	// run first strategy
	strategy := &firstAlgorithm{}
	pv := initStrategy(strategy)
	pv.userFilters["role"] = 1
	getData(pv)

	// run second strategy
	strategySecond := &secondAlgorithm{}
	pv.userFilters["pop"] = 2
	pv.setStrategy(strategySecond)
	getData(pv)

	// мы можем управлять вызовом разных алгоритмов в зависимости от контекста и пользовательских фильтров
}
