package main

import (
	"fmt"
	"sort"
	"strings"
)

//Написать функцию поиска всех множеств анаграмм по словарю.

//Например:
//'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
//'листок', 'слиток' и 'столик' - другому.

//Требования:
//Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
//Выходные данные: ссылка на мапу множеств анаграмм
//Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
//слово из множества.
//Массив должен быть отсортирован по возрастанию.
//Множества из одного элемента не должны попасть в результат.
//Все слова должны быть приведены к нижнему регистру.
//В результате каждое слово должно встречаться только один раз.

type R []rune

func (s R) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s R) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s R) Len() int {
	return len(s)
}

func IndividualAndSort(str []string) []string {
	M := make(map[string][]bool)
	var answer []string
	for _, val := range str {
		M[val] = append(M[val], true)
	}
	for key, _ := range M {
		answer = append(answer, key)
	}
	sort.Strings(answer)
	//fmt.Println(answer)
	return answer

}

func main() {
	JustMap := make(map[string][]string)
	AnswerMap := make(map[string][]string)

	Words := []string{
		"Пятак",
		"пятка",
		"Тяпка",
		"слиток",
		"слиток",
		"столик",
		"листок",
		"Топот",
		"Пятак",
		"Потоп",
		"Листок",
	}
	//Группируем слова, ключ которых является отсортированная строка
	for _, val := range Words {
		LowWord := strings.ToLower(val)
		Rn := R(LowWord)
		sort.Sort(Rn)
		JustMap[string(Rn)] = append(JustMap[string(Rn)], LowWord)

	}
	//Сортируем данные и кладем результат в новую мапу
	for key, v := range JustMap {
		NewV := IndividualAndSort(v)
		key = NewV[0]
		NewV = append(NewV[:0], NewV[1:]...)
		AnswerMap[key] = NewV

	}
	fmt.Println(AnswerMap)
}
