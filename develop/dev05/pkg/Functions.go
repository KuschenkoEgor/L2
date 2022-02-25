package pkg

import (
	"bufio"
	"io"
	"strings"
)

//Чтение файла и получение строк
func ReadData(file io.Reader) ([]string, error) {
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

//Вывод строк, содержащие определенный абзац
func JustSearch(lines []string, SearchString string) []string {
	Result := make([]string, 0)
	for _, line := range lines {
		if strings.Contains(line, SearchString) {
			Result = append(Result, line)
		}
	}
	return Result
}

//Поиск индекса строк, в которых есть совпадение
func SearchIndex(lines []string, SearchString string) []int {
	Index := make([]int, 0)
	for i, line := range lines {
		if strings.Contains(line, SearchString) {
			Index = append(Index, i)
		}
	}
	return Index
}

//Поиск +N строк после совпадения
func After(lines []string, SearchString string, N int) []string {
	Result := make([]string, 0)
	Index := SearchIndex(lines, SearchString)
	for _, idx := range Index {
		for i := idx; i < idx+N+1; i++ {
			Result = append(Result, lines[i])
		}
	}
	return Result
}

//Поиск -N строк до совпадения
func Before(lines []string, SearchString string, N int) []string {
	Result := make([]string, 0)
	Index := SearchIndex(lines, SearchString)
	for _, idx := range Index {
		for i := idx; i < idx-N-1; i-- {
			Result = append(Result, lines[i])
		}
	}
	return Result
}

//Печатать строки вокруг найденной
func Context(lines []string, SearchString string, N int) []string {
	Result := make([]string, 0)
	Index := SearchIndex(lines, SearchString)
	for _, idx := range Index {
		for i := idx - N - 1; i < idx+N+1; i++ {
			Result = append(Result, lines[i])
		}
	}
	return Result
}

//Поиск количества строк
func Count(lines []string, SearchString string) int {
	Index := SearchIndex(lines, SearchString)

	return len(Index)
}

//Игнорировать регистр
func IgnoreCase(lines []string, SearchString string) []string {
	LowString := strings.ToLower(SearchString)
	Result := make([]string, 0)
	for _, line := range lines {
		LowLine := strings.ToLower(line)
		if strings.Contains(LowLine, LowString) {
			Result = append(Result, line)
		}
	}
	return Result
}

//Исключение строк с совпадениями
func Invert(lines []string, SearchString string) []string {
	for i, line := range lines {
		if strings.Contains(line, SearchString) {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}
	return lines
}

//Печать номера строки
func LineNum(lines []string, SearchString string) []int {
	Result := make([]int, 0)
	for i, line := range lines {
		if strings.Contains(line, SearchString) {
			Result = append(Result, i+1)
		}
	}
	return Result
}
