package main

import "GolandProjects/L2/pattern/command/pkg"

func main() {
	tv := &pkg.Tv{}

	OnCommand := &pkg.OnCommand{
		Device: tv,
	}

	OffCommand := &pkg.OffCommand{
		Device: tv,
	}

	onButton := &pkg.Button{
		Command: OnCommand,
	}
	onButton.Press()
	offButton := &pkg.Button{
		Command: OffCommand,
	}
	offButton.Press()
}

//Преимущества:
//Убирает прямую зависимость между объектами, вызывающими операции,
//и объектами, которые их непосредственно выполняют.

//Позволяет реализовать простую отмену и повтор операций.
//Позволяет реализовать отложенный запуск операций.
//Позволяет собирать сложные команды из простых.
//Реализует принцип открытости/закрытости.

//Недостатки:
//Усложняет код программы из-за введения множества дополнительных классов.
