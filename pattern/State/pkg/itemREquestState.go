package pkg

import "fmt"

type ItemRequestState struct {
	VendingMachine *VendingMachine
}

func (i *ItemRequestState) RequestItem() error{
	return fmt.Errorf("Item already requested")
}

func (i *ItemRequestState) AddItem(count int) error{
	return fmt.Errorf("Item Dispense in progress")
}

func (i *ItemRequestState) InsertMoney(money int) error{
	if money < i.VendingMachine.ItemPrice{
		fmt.Errorf("Inserted money is less. Please insert %d", i.VendingMachine.ItemPrice)
	}
	fmt.Println("Money entered")
	i.VendingMachine.SetState(i.VendingMachine.HasMoney)
	return nil
}

func (i *ItemRequestState) DispenseItem() error{
	return fmt.Errorf("Please insert money")
}

