package main

import "fmt"

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("Item already requested")
}

func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("Please wait, we are already giving you an item")
}
func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}

func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		fmt.Printf("You inserted %d cents, item price is %d cents \n", money, i.vendingMachine.itemPrice)
		return fmt.Errorf("Not enough money")
	}
	fmt.Printf("You inserted %d cents \n", money)
	i.vendingMachine.setCurrentState(i.vendingMachine.hasMoney)
	return nil
}
