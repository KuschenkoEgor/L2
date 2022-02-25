package pkg

import "strings"

func ChoiceColumn(lines []string, N int, Tab string) []string {

	Result := make([]string, 0)
	for _, line := range lines {
		x := strings.Split(line, Tab)
		Result = append(Result, x[N-1])
	}
	return Result
}

func Separated(lines []string, Tab string) []string {
	Result := make([]string, 0)
	for _, line := range lines {
		if strings.Contains(line, Tab) {
			Result = append(Result, line)
		}
	}
	return Result
}
