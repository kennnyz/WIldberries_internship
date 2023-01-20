package main

import (
	"errors"
	"fmt"
	"time"
)

type Shop struct {
	Name     string
	Products []Product
}

func (shop Shop) Sell(user User, product string) error {
	println("[Магазин] Запрос к полльзователю, для получения остатка по карте")
	time.Sleep(500 * time.Millisecond)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] Проверка - может ли [%s] купить товар [%s]", user.Name, product)
	time.Sleep(500 * time.Millisecond)

	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] Не достаточно средств! ")
		}
		fmt.Printf("[Магазин] Товар [%s] - куплен!", product)
	}
	return nil
}
