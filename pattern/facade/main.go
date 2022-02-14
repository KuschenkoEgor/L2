//Фасад позволяет представить простой или
//урезанный интерфейс к сложной подсистеме.

//Преимущество:  Изолирует клиентов от компонентов сложной подсистемы.
//Недостаток: Фасад рискует стать божественным объектом, привязанным ко всем классам программы
package main

import (
	"GolandProjects/L2/pattern/facade/pkg"
	"log"
)

func main() {

	x := pkg.NewWalletFacade("Admin", 1234)

	err := x.AddMoneyToWallet("Admin", 1234, 1000)

	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	err = x.DebitMoneyFromWallet("Admin", 1234, 500)

	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

}
