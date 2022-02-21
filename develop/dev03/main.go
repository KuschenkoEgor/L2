package main

import (
	"GolandProjects/L2_WB/develop/dev03/pkg"
	"flag"
	"fmt"
	"os"
	"sort"
)

//Отсортировать строки в файле по аналогии с консольной утилитой sort
//(man sort — смотрим описание и основные параметры): на входе подается
//файл из несортированными строками, на выходе — файл с отсортированными.

//Реализовать поддержку утилитой следующих ключей:
//
//-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
//-n — сортировать по числовому значению
//-r — сортировать в обратном порядке
//-u — не выводить повторяющиеся строки

//Дополнительно

//Реализовать поддержку утилитой следующих ключей:
//
//-M — сортировать по названию месяца
//-b — игнорировать хвостовые пробелы
//-c — проверять отсортированы ли данные
//-h — сортировать по числовому значению с учетом суффиксов

type Arguments struct {
	k int
	n bool
	r bool
	u bool
	M bool
	b bool
	c bool
	h bool
}

var err error

func FullSort(lines []string, Arg Arguments) []string {
	var Result []string
	switch {
	case Arg.k > 0:
		Result = pkg.SortByColumn(Arg.k, lines)

	case Arg.n:
		Result, err = pkg.NumberSort(lines)
		if err != nil {
			fmt.Println(err)
		}

	case Arg.r:
		Result = pkg.ReverseSort(lines)

	case Arg.u:
		Result = pkg.DuplicateFilter(lines)

	case Arg.M:
		Result = pkg.MonthSort(lines)

	case Arg.c:
		Check := pkg.CheckSort(lines)
		if Check {
			fmt.Println("данные отсортированы")
		} else {
			fmt.Println("Данные не отсортированы")
		}

	case Arg.h:
		Result = pkg.NumberSortWithSuffix(lines)

	default:
		sort.Strings(lines)
		Result = lines
	}

	return Result
}

func main() {
	var Arg Arguments
	lines := make([]string, 0)

	flag.IntVar(&Arg.k, "k", 0, "Сортировка по определенному столбцу")
	flag.BoolVar(&Arg.n, "n", false, "Сортировка по числовому значению")
	flag.BoolVar(&Arg.r, "r", false, "Сортировка в обратном порядке")
	flag.BoolVar(&Arg.u, "u", false, "Не выводить повторяющиеся строки")
	flag.BoolVar(&Arg.M, "M", false, "Сортировка по названию месяца")
	flag.BoolVar(&Arg.b, "b", false, "Игнор хвостовых пробелов")
	flag.BoolVar(&Arg.c, "c", false, "Проверка сортировки")
	flag.BoolVar(&Arg.h, "h", false, "Сортировка по числовому значению с учетом суффиксов")

	flag.Parse()

	file, _ := os.Open("/home/zhora/GolandProjects/L2_WB/develop/dev03/testM.txt")

	lines, err = pkg.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	Answer := FullSort(lines, Arg)
	fmt.Println(Answer)
}
