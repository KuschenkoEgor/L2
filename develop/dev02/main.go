package main

//Создать Go-функцию, осуществляющую примитивную
//распаковку строки, содержащую повторяющиеся символы/руны,
//например:
//"a4bc2d5e" => "aaaabccddddde"
//"abcd" => "abcd"
//"45" => "" (некорректная строка)
//"" => ""

//Реализовать поддержку escape-последовательностей.
//Например:
//qwe\4\5 => qwe45 (*)
//qwe\45 => qwe44444 (*)
//qwe\\5 => qwe\\\\\ (*)

//В случае если была передана некорректная строка,
//функция должна возвращать ошибку. Написать unit-тесты.
import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ConvertString(rs, answer []rune, k, ln int) ([]rune, error) {
	for i := 0; i < ln; i++ {

		if unicode.IsDigit(rs[0]) {
			break
		}

		if unicode.IsLetter(rs[i]) {
			answer = append(answer, rs[i])

		} else if unicode.IsDigit(rs[i]) && i != 0 && unicode.IsLetter(rs[i-1]) {
			x, _ := strconv.Atoi(string(rs[i]))
			repeat := strings.Repeat(string(rs[i-1]), x-1)
			rp := []rune(repeat)
			answer = append(answer, rp...)

		} else if i != 0 && strings.Contains(string(rs[i-1]), `\`) && strings.Contains(string(rs[i-2]), `\`) != true {
			answer = append(answer, rs[i])

		} else if unicode.IsDigit(rs[i]) && i != 0 && strings.Contains(string(rs[i-2]), `\`) {
			x, _ := strconv.Atoi(string(rs[i]))
			repeat := strings.Repeat(string(rs[i-1]), x-1)
			rp := []rune(repeat)
			answer = append(answer, rp...)
		}

	}
	for _ = range answer {
		k++
	}
	if k == 0 {
		err := errors.New("Некорректная строка")
		return nil, err
	}

	return answer, nil
}

func main() {
	var k int
	var Str string
	var err error
	fmt.Println("Введите строку:")
	fmt.Scan(&Str)
	rs := []rune(Str)
	answer := make([]rune, 0)
	ln := utf8.RuneCountInString(Str)

	answer, err = ConvertString(rs, answer, k, ln)
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println("Отформатированная строка:", string(answer))
}
