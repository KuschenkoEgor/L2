package main

import (
	"GolandProjects/L2_WB/develop/dev08/pkg"
	"fmt"
	"os"
	"strings"
)

//Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
//
//- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
//- pwd - показать путь до текущего каталога
//- echo <args> - вывод аргумента в STDOUT
//- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
//- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

//Так же требуется поддерживать функционал fork/exec-команд

//Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

func GoShell() error {

	for {
		line, err := pkg.ReadCmdLine()
		if err != nil {
			return err
		}

		if line == `\quit` {
			break
		}

		cmd := strings.Split(line, "|")
		for _, c := range cmd {

			R, err := pkg.AllCommands(c)

			fmt.Println(R)

			if err != nil {
				return err
			}
		}

	}

	return nil
}

func main() {

	err := GoShell()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
