package main

import (
	"fmt"
)

var (
	bank = Bank{
		Name:  "Bank",
		Cards: []Card{},
	}
	card1 = Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = Card{
		Name:    "CRD-2",
		Balance: -20,
		Bank:    &bank,
	}

	user = User{
		Name: "User 1",
		Card: &card1,
	}

	user2 = User{
		Name: "User 2",
		Card: &card2,
	}

	prod = Product{
		Name:  "Cheese",
		Price: 150,
	}

	shop = Shop{
		Name: "Magnet",
		Products: []Product{
			prod,
		},
	}
)

func main() {
	println("[Банк] Выпуск карт")
	bank.Cards = append(bank.Cards, card1, card2)

	fmt.Printf("[%s]", user.Name)
	err := shop.Sell(user, prod.Name)
	if err != nil {
		println(err)
		return
	}

	fmt.Printf("[%s]\n", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}

}
