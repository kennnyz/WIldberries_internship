package main

import (
	"fmt"
	"sort"
	"strings"
)

//Написать функцию поиска всех множеств анаграмм по словарю.
//
//
//Например:
//'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
//'листок', 'слиток' и 'столик' - другому.
//
//
//Требования:
//Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
//Выходные данные: ссылка на мапу множеств анаграмм
//Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
//слово из множества.
//Массив должен быть отсортирован по возрастанию.
//Множества из одного элемента не должны попасть в результат.
//Все слова должны быть приведены к нижнему регистру.
//В результате каждое слово должно встречаться только один раз.

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.ToLower(strings.Join(s, ""))
}

func Pop(pop *[7]string) *map[string][]string {

	pp := []string{}
	for i := range pop {
		pp = append(pp, strings.ToLower(pop[i]))
	}

	ans := map[string][]string{}
	uniquie := make(map[int]int)

	for i := 0; i < len(pop)-1; i++ {

		if uniquie[i] == 1 {
			continue
		}

		uniquie[i] = 1

		// Сортирую слово и перевожу все символы в нижний регистр
		word := sortString(pp[i])
		anagrams := []string{}

		anagrams = append(anagrams, strings.ToLower(pp[i]))
		for j := i + 1; j < len(pp); j++ {
			pare := sortString(pp[j])
			// чек если встретился кейс где слово является анаграммой и при этом это слово не принадлежит множеству
			if word == pare && uniquie[j] != 1 {
				anagrams = append(anagrams, strings.ToLower(pp[j]))
				uniquie[j] = 1
			} else {
				continue
			}
		}
		ans[strings.ToLower(pp[i])] = anagrams
	}

	return &ans
}

func main() {
	myDictr := [7]string{"Изба", "бАзи", "пЯткА", "тЯпка", "слиток", "столик", "листок"}
	jjoj := Pop(&myDictr)
	fmt.Println(jjoj)
}
