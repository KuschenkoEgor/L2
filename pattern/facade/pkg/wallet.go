package pkg

import "fmt"

type wallet struct {
	balance int
}

func NewWallet() *wallet {
	return &wallet{balance: 0}
}

func (w *wallet) CreditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}
func (w *wallet) DebitBalance(amount int) error {
	if amount > w.balance {
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance -= amount
	return nil
}
