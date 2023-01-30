package main

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (n *NoItemState) addItem(count int) error {
	n.vendingMachine.IncrementItemCount(count)
	n.vendingMachine.setCurrentState(n.vendingMachine.hasItem)
	return nil
}

func (n *NoItemState) requestItem() error {
	return fmt.Errorf("No item in the machine")
}

func (n *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("No item in the machine")
}

func (n *NoItemState) dispenseItem() error {
	return fmt.Errorf("No item in the machine")
}
