package main

import "fmt"

type hasMoneyState struct {
	vendingMachine *VendingMachine
}

func (h *hasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (h *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (h *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (i *hasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setCurrentState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setCurrentState(i.vendingMachine.hasItem)
	}
	return nil
}
