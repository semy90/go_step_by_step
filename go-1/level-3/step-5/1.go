package main

/*
Разработайте программу, которая будет анализировать текст и предоставлять по нему статистические данные. Для этого напишите функцию AnalyzeText(text string), которая принимает на вход текст и печатает в консоль такой результат анализа:

Количество слов: *целое число*
Количество уникальных слов: *целое число*
Самое часто встречающееся слово: "*строка*" (встречается *целое число* раз)
Топ-5 самых часто встречающихся слов:
"*строка*": *целое число* раз
"*строка*": *целое число* раз
"*строка*": *целое число* раз
"*строка*": *целое число* раз
"*строка*": *целое число* раз
Нужно вычислить:

Количество слов в тексте (считаем, что все слова в тексте разделяются только пробелами, точками, запятыми, восклицательным или вопросительным знаками)
Количество уникальных слов (регистр не влияет на уникальность слова, то есть Привет и привет считаются одним словом)
Слово, которое встретилось в тексте больше всего раз (без учёта регистра)
Топ-5 самых часто встречающихся слов, по убыванию. Регистр также учитывать не нужно. Гарантируется, что в тексте есть хотя бы пять уникальных слов и нет двух (и более) слов в топе, которые встречаются одинаковое количество раз.
Для подсчёта топа самых популярных слов напишите вспомогательную функцию func getTopWords(wordMap map[string]int, n int) []string, которая принимает мапу с ключом слов и числом, сколько раз встречается это слово, а также числом n — количеством позиций в топе. Функция должна вернуть слайс с топ-n словами.
*/
import (
	"fmt"
	"sort"
	"strings"
)

type wordcount struct {
	Word  string
	Count int
}

func getTopWords(wordMap map[string]int, n int) []string {
	var sliceOfWord []wordcount

	for word, count := range wordMap {
		sliceOfWord = append(sliceOfWord, wordcount{word, count})
	}

	sort.Slice(sliceOfWord, func(i int, j int) bool {
		return sliceOfWord[i].Count > sliceOfWord[j].Count
	})

	var ans []string

	for i := range sliceOfWord {
		ans = append(ans, sliceOfWord[i].Word)
	}

	return ans[:n]
}

func AnalyzeText(text string) {
	text = strings.Replace(text, "!", " ", -1)
	text = strings.Replace(text, "?", " ", -1)
	text = strings.Replace(text, ".", " ", -1)
	text = strings.Replace(text, ",", " ", -1)
	text = strings.ToLower(text)
	textArray := strings.Fields(text)
	wordMap := map[string]int{}

	for i := range textArray {
		s := textArray[i]
		if _, f := wordMap[s]; !f {
			wordMap[s] = 1
		} else {
			wordMap[s]++
		}
	}

	topWords := getTopWords(wordMap, 5)

	fmt.Printf("Количество слов: %d\n", len(textArray))
	fmt.Printf("Количество уникальных слов: %d\n", len(wordMap))
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", topWords[0], wordMap[topWords[0]])
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for _, word := range topWords {
		fmt.Printf("\"%s\": %d раз\n", word, wordMap[word])
	}

}
