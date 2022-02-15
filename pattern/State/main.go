package main

import (
	"example.com/m/v2/L2_WB/pattern/State/pkg"
	"log"
)

func main() {
	vendingVachine := pkg.NewVendingMachine(1,10)

	err := vendingVachine.RequestItem()
	if err != nil{
		log.Fatal(err.Error())
	}

	err = vendingVachine.DispenseItem()
	if err != nil{
		log.Fatal(err.Error())
	}

	err = vendingVachine.AddItem(2)
	if err != nil{
		log.Fatal(err.Error())
	}

	err = vendingVachine.RequestItem()
	if err != nil{
		log.Fatal(err.Error())
	}

	err = vendingVachine.InsertMoney(10)
	if err != nil{
		log.Fatal(err.Error())
	}

	err = vendingVachine.DispenseItem()
	if err != nil{
		log.Fatal(err.Error())
	}
}
