package main

import (
	"GolandProjects/L2_WB/develop/dev05/pkg"
	"flag"
	"fmt"
	"os"
	"strconv"
)

//Реализовать утилиту фильтрации по аналогии с консольной
//утилитой (man grep — смотрим описание и основные параметры).

//Реализовать поддержку утилитой следующих ключей:
//-A - "after" печатать +N строк после совпадения
//-B - "before" печатать +N строк до совпадения
//-C - "context" (A+B) печатать ±N строк вокруг совпадения
//-c - "count" (количество строк)
//-i - "ignore-case" (игнорировать регистр)
//-v - "invert" (вместо совпадения, исключать)
//-F - "fixed", точное совпадение со строкой, не паттерн
//-n - "line num", напечатать номер строки

type Args struct {
	A       int
	B       int
	C       int
	c       bool
	i       bool
	v       bool
	F       bool
	n       bool
	SrchStr string
}

var Ar Args
var SearS string

var err error

func Search(lines []string, A Args) []string {
	Answer := make([]string, 0)

	switch {
	case A.A > 0:
		Answer = pkg.After(lines, A.SrchStr, A.A)
	case A.B > 0:
		Answer = pkg.Before(lines, A.SrchStr, A.B)
	case A.C > 0:
		Answer = pkg.Context(lines, A.SrchStr, A.C)
	case A.c:
		N := pkg.Count(lines, A.SrchStr)
		Answer = append(Answer, strconv.Itoa(N))
	case A.i:
		Answer = pkg.IgnoreCase(lines, A.SrchStr)
	case A.v:
		Answer = pkg.Invert(lines, A.SrchStr)
	case A.F:
		Answer = pkg.JustSearch(lines, A.SrchStr)
	case A.n:
		X := pkg.LineNum(lines, A.SrchStr)
		for _, val := range X {
			Answer = append(Answer, strconv.Itoa(val))
		}
	default:
		Answer = pkg.JustSearch(lines, A.SrchStr)
	}
	return Answer
}

func main() {

	lines := make([]string, 0)

	flag.IntVar(&Ar.A, "A", 0, "Печать +N строк после совпадения")
	flag.IntVar(&Ar.B, "B", 0, "Печать +N строк до совпадения")
	flag.IntVar(&Ar.C, "C", 0, "Печать +N строк до и после совпадения")
	flag.BoolVar(&Ar.c, "c", false, "Печать количества совпаденных строк")
	flag.BoolVar(&Ar.i, "i", false, "Игнорирование регистра")
	flag.BoolVar(&Ar.v, "v", false, "Исключать строки с совпадениями")
	flag.BoolVar(&Ar.F, "F", false, "Точное совпадение со строкой")
	flag.BoolVar(&Ar.n, "n", false, "Печать номера строки")

	flag.Parse()

	Ar.SrchStr = flag.Args()[0]

	file, _ := os.Open("/home/zhora/GolandProjects/L2_WB/develop/dev05/test.txt")

	lines, err = pkg.ReadData(file)
	if err != nil {
		fmt.Println(err)
	}

	Result := Search(lines, Ar)
	fmt.Println(Result)

}
