package pkg

import "fmt"

type account struct {
	Name string
}

func NewAccount(accountName string) *account {
	return &account{Name: accountName}
}

func (a *account) checkAccount(accountName string) error {
	if a.Name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}
