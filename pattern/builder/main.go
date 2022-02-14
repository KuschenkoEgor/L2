package main

import (
	"GolandProjects/L2/pattern/builder/pkg"
	"fmt"
)

//Паттерн Строитель позволяет собирать объекты пошагово,
//вызывая только те шаги, которые вам нужны. А значит,
//больше не нужно пытаться «запихнуть» в конструктор все возможные опции продукта.
func main() {

	normalBuilder := pkg.GetBuilder("normal")
	iglooBuilder := pkg.GetBuilder("igloo")

	director := pkg.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.DoorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)
}

//Преимущества:
//Позволяет создавать продукты пошагово.
// Позволяет использовать один и тот же код для создания различных продуктов.
// Изолирует сложный код сборки продукта от его основной бизнес-логики.

//Недостатки:
//Усложняет код программы из-за введения дополнительных классов.
// Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.
