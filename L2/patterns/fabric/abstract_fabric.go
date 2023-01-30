package fabric

import "fmt"

//Абстрактная фабрика — это порождающий паттерн проектирования,
//который решает проблему создания целых семейств связанных продуктов, без указания конкретных классов продуктов.

//Абстрактная фабрика задаёт интерфейс создания всех доступных типов продуктов, а каждая конкретная реализация фабрики порождает продукты одной из вариаций.
//Клиентский код вызывает методы фабрики для получения продуктов, вместо самостоятельного создания с помощью оператора new.
//При этом фабрика сама следит за тем, чтобы создать продукт нужной вариации.

type iPizzaFactory interface {
	makePizza() Ipizza
	makePotato() Ipotato
}

func GetPizzaFactory(brand string) iPizzaFactory {
	if brand == "dodo" {
		return &Dodo{}
	}

	if brand == "domino" {
		return &Domino{}
	}
	fmt.Println("Wrong brand type passed")
	return nil

}

type Dodo struct {
}

func (d *Dodo) makePizza() Ipizza {
	return &DodoPizza{
		pizza{
			name:  "DodoPizza with CocaCola",
			price: 400,
		},
	}
}

func (d *Dodo) makePotato() Ipotato {
	return &DominoPotato{
		potato{
			name:  "DodoPotato ",
			price: 150,
		},
	}
}

type Domino struct {
}

func (d *Domino) makePizza() Ipizza {
	return &DominoPizza{
		pizza{
			name:  "DominoPizza with HachaPuri",
			price: 389,
		},
	}
}

func (d *Domino) makePotato() Ipotato {
	return &DominoPotato{
		potato{
			name:  "Domino Potato ",
			price: 169,
		},
	}
}

type Ipizza interface {
	setName(name string)
	setPrice(price int)
	getName() string
	getPrice() int
}

type pizza struct {
	name  string
	price int
}

func (p *pizza) setName(name string) {
	p.name = name
}

func (p *pizza) setPrice(price int) {
	p.price = price
}

func (p *pizza) getName() string {
	return p.name
}

func (p *pizza) getPrice() int {
	return p.price
}

type DodoPizza struct {
	pizza
}

type DominoPizza struct {
	pizza
}

type Ipotato interface {
	setName(name string)
	setPrice(price int)
	getName() string
	getPrice() int
}

type potato struct {
	name  string
	price int
}

func (p *potato) setName(name string) {
	p.name = name
}
func (p *potato) setPrice(price int) {
	p.price = price
}
func (p *potato) getName() string {
	return p.name
}
func (p *potato) getPrice() int {
	return p.price
}

type DodoPotato struct {
	potato
}

type DominoPotato struct {
	potato
}

func main() {
	dodoFactory := GetPizzaFactory("dodo")

	pizza := dodoFactory.makePizza()
	fmt.Println(pizza.getName())
}
