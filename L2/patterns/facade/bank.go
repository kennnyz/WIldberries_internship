package main

import (
	"errors"
	"fmt"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

func (bank Bank) CheckBalance(cardNumber string) error {
	println(fmt.Sprintf("[Банк] получения остатка по карте %s", cardNumber))
	time.Sleep(time.Millisecond * 300)
	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] недостаточно средств! ")
		}
	}
	println("[Банк] Остаток положительный!")
	return nil
}
