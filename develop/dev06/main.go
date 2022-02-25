package main

import (
	"GolandProjects/L2_WB/develop/dev06/pkg"
	"flag"
	"fmt"
)

//Реализовать утилиту аналог консольной команды cut (man cut).
//Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB)
//на колонки и выводить запрошенные.

//Реализовать поддержку утилитой следующих ключей:
//-f - "fields" - выбрать поля (колонки)
//-d - "delimiter" - использовать другой разделитель
//-s - "separated" - только строки с разделителем

type Args struct {
	f int
	d string
	s bool
}

func Cut(lines []string, A Args) []string {
	var Answer []string
	switch {
	case A.s || A.d != " ":

		Answer = pkg.Separated(lines, A.d)

	case A.f > 0:
		Answer = pkg.ChoiceColumn(lines, A.f, A.d)

	default:
		fmt.Println("Ошибка! Укажите флаги")

	}
	return Answer
}

func main() {

	var A Args

	flag.IntVar(&A.f, "f", 0, "Выбор определенной колонки")
	flag.StringVar(&A.d, "d", " ", "Использование другого разделителя")
	flag.BoolVar(&A.s, "s", false, "Выбор строк только с разделителем")

	flag.Parse()

	Result := Cut(flag.Args(), A)

	fmt.Println(Result)
}
