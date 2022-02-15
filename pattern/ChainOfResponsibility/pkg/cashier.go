package pkg

import "fmt"

type Cashier struct {
	Next Departament
}

func (c *Cashier) Execute(p *Patient){
	if p.PaymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient")
}

func (c *Cashier) SetNext(next Departament){
	c.Next = next
}
