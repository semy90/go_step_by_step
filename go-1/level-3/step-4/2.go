package main

/*
В самом большом тайнике, ключ от которого вы помогли найти Арсению, оказалась целая гора NFT-артефактов. Помогите ему ещё раз и найдите общую сумму этих артефактов.

Напишите функцию SumOfValuesInMap(m map[int]int) int, которая возвращает сумму значений в мапе.
*/
func SumOfValuesInMap(m map[int]int) int {
	s := 0
	for k := range m {
		s += m[k]
	}
	return s
}
