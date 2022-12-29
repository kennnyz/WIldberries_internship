package main

import "fmt"

type ITransport interface {
	Drive()
}

type Car struct {
}

func (c *Car) Drive() {
	fmt.Println("Путешественник на машине")
}

// Структура к которой нужен адаптер

type Camel struct {
}

type CamelToTransportAdapter struct {
	camel *Camel
}

func (adapter *CamelToTransportAdapter) Drive() {
	fmt.Println("Верблюд идет по пескам пустыни")
}
func going(d ITransport) {
	d.Drive()
}
func main() {
	car := &Car{}
	camel := &Camel{}
	adapter := &CamelToTransportAdapter{camel: camel}
	going(car)
	going(adapter)
}
