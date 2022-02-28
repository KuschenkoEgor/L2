package main

import (
	"testing"
	"unicode/utf8"
)

func TestConvertString(t *testing.T) {
	tables := []struct {
		s   string
		res string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{`qwe\4\5`, "qwe45"},
		{`qwe\45`, "qwe44444"},
		{`qwe\\5`, `qwe\\\\\`},
	}
	for _, table := range tables {
		var k int
		ln := utf8.RuneCountInString(table.s)
		rs := []rune(table.s)
		answer := make([]rune, 0)
		TotalRn, _ := ConvertString(rs, answer, k, ln)

		if string(TotalRn) != table.res {
			t.Errorf("Incorrect result. Expect %v, got %v", table.res, string(TotalRn))
		}

	}
}
