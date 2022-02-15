package pkg

type State interface {
	AddItem(int) error
	RequestItem() error
	InsertMoney(Money int) error
	DispenseItem() error
}
