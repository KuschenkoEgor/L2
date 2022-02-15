package main

import "example.com/m/v2/L2_WB/pattern/ChainOfResponsibility/pkg"

func main() {
	cashier := &pkg.Cashier{}

	medical := &pkg.Medical{}
	medical.SetNext(cashier)

	doctor := &pkg.Doctor{}
	doctor.SetNext(medical)

	reception :=&pkg.Reception{}
	reception.SetNext(doctor)

	patient := &pkg.Patient{Name: "Egor"}
	reception.Execute(patient)

}

//Применимость:
//Когда программа должна обрабатывать разнообразные запросы несколькими
//способами, но заранее неизвестно, какие конкретно запросы будут приходить
//и какие обработчики для них понадобятся.


//Достоинства:
//Уменьшает зависимость между клиентом и обработчиками.
//Реализует принцип единственной обязанности.
//Реализует принцип открытости/закрытости.


//Недостатки:
// Запрос может остаться никем не обработанным.Ы
