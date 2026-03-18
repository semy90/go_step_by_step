package main

/*
Как вы уже знаете, в телефоне Арсения очень много контактов. Он решил найти все повторяющиеся и подсчитать их. Помогите ему найти дубликаты.

Для этого напишите функцию CountingSort(contacts []string) map[string]int, которая принимает слайс строк и возвращает мапу, где ключ — это элемент слайса, а значение — количество его повторов.
*/
func CountingSort(contacts []string) map[string]int {
	ans := map[string]int{}
	for i := range contacts {
		if _, f := ans[contacts[i]]; !f {
			ans[contacts[i]] = 1
		} else {
			ans[contacts[i]]++
		}
	}
	return ans
}
