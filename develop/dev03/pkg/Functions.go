package pkg

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

//Чтение файла и получение строк
func ReadFile(file io.Reader) ([]string, error) {
	var lines []string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		lines = append(lines, line)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return lines, nil
}

//Удаление дублированных строк
func DuplicateFilter(lines []string) []string {
	Answer := make([]string, 0)
	Map := make(map[string][]bool)

	for _, line := range lines {
		Map[line] = append(Map[line], true)
	}
	for key, _ := range Map {
		Answer = append(Answer, key)
	}

	return Answer
}

//Разбиение строки на отдельные слова
func WordsFromString(numColumn int, line string) string {
	if line == "" {
		return ""
	}
	Words := strings.Fields(line)
	ln := utf8.RuneCountInString(line)
	if numColumn < ln {
		return Words[numColumn]
	} else {
		return ""
	}

}

//Сортировка по отдельной колонке
func SortByColumn(num int, lines []string) []string {
	ReceivedWords := make([]string, 0)
	Result := make([]string, 0)
	for _, line := range lines {
		//Получение слова из определенного столбца
		TheRightWord := WordsFromString(num, line)
		ReceivedWords = append(ReceivedWords, TheRightWord)
	}
	sort.Strings(ReceivedWords)
	for _, w := range ReceivedWords {
		for _, line := range lines {
			if strings.Contains(line, w) {
				Result = append(Result, line)
			}
		}
	}
	return Result
}

//Сортировка по месяцам
func MonthSort(lines []string) []string {
	var Months = map[string]int{
		"JAN": 1,
		"FAB": 2,
		"MAR": 3,
		"APR": 4,
		"MAY": 5,
		"JUN": 6,
		"JUL": 7,
		"AUG": 8,
		"SEP": 9,
		"OCT": 10,
		"NOV": 11,
		"DEC": 12,
	}
	NotMoth := make([]string, 0)
	AnswerInt := make([]int, 0)
	Result := make([]string, 0)
	for _, month := range lines {
		if val, ok := Months[month]; ok {
			AnswerInt = append(AnswerInt, val)
		} else {
			NotMoth = append(NotMoth, month)
		}

	}
	sort.Ints(AnswerInt)
	for _, num := range AnswerInt {
		for key, _ := range Months {
			if Months[key] == num {
				Result = append(Result, key)
			}
		}
	}
	for _, v := range NotMoth {
		Result = append(Result, v)
	}
	return Result
}

//Сортировка по числам
func NumberSort(lines []string) ([]string, error) {
	Nums := make([]int, 0)
	Result := make([]string, 0)
	for _, nums := range lines {
		n, err := strconv.Atoi(nums)
		if err != nil {
			return nil, err
		}
		Nums = append(Nums, n)
	}
	sort.Ints(Nums)
	for _, val := range Nums {
		Ns := strconv.Itoa(val)
		Result = append(Result, Ns)
	}
	return Result, nil
}

//Сортировка по числовому значению с учетом суффиксов
func NumberSortWithSuffix(Lines []string) []string {
	sort.Strings(Lines)
	return Lines
}

//Проверка сортировки
func CheckSort(lines []string) bool {
	Answer := sort.StringsAreSorted(lines)
	return Answer
}

//Сортировка в обратном порядке
func ReverseSort(lines []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(lines)))
	return lines
}
