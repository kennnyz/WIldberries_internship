package main

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (h *HasItemState) addItem(count int) error {
	fmt.Printf("Adding %d items \n", count)
	h.vendingMachine.IncrementItemCount(count)
	return nil
}

func (h *HasItemState) requestItem() error {
	if h.vendingMachine.itemCount <= 0 {
		h.vendingMachine.setCurrentState(h.vendingMachine.itemRequested)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requestd\n")
	h.vendingMachine.setCurrentState(h.vendingMachine.itemRequested)
	return nil
}

func (h *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("Please request an item first")
}

func (h *HasItemState) dispenseItem() error {
	return fmt.Errorf("Please request an item first")
}
