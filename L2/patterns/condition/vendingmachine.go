package main

import "fmt"

// конетекст

//hasItem (имеетПредмет)
//noItem (неИмеетПредмет)
//itemRequested (выдаётПредмет)
//hasMoney (получилДеньги)

type VendingMachine struct {
	hasItem       State
	itemRequested State
	noItem        State
	hasMoney      State

	currentState State

	itemPrice int
	itemCount int
}

func NewVendingMachine(itemPrice, itemCount int) *VendingMachine {
	vendingMachine := &VendingMachine{
		itemPrice: itemPrice,
		itemCount: itemCount,
	}

	vendingMachine.hasItem = &HasItemState{vendingMachine: vendingMachine}
	vendingMachine.itemRequested = &ItemRequestedState{vendingMachine: vendingMachine}
	vendingMachine.noItem = &NoItemState{vendingMachine: vendingMachine}
	vendingMachine.hasMoney = &hasMoneyState{vendingMachine: vendingMachine}

	if vendingMachine.itemCount > 0 {
		vendingMachine.setCurrentState(vendingMachine.hasItem)
	} else {
		vendingMachine.setCurrentState(vendingMachine.noItem)
	}

	return vendingMachine
}

func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *VendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *VendingMachine) setCurrentState(state State) {
	v.currentState = state
}

func (v *VendingMachine) IncrementItemCount(count int) {
	fmt.Printf("Adding %d items \n", count)
	v.itemCount += count
}
