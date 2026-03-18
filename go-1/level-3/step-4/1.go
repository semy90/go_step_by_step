package main

/*
Арсению в руки попал набор ключей–кодов от виртуальных тайников и их номера. Помогите ему найти самый большой ключ — ведь чем он больше, тем больше тайник!

Напишите функцию FindMaxKey(m map[int]int) int, которая принимает мапу и возвращает значение наибольшего ключа.
*/
import "math"

func FindMaxKey(m map[int]int) int {
	mx := math.MinInt
	for k := range m {
		if mx < k {
			mx = k
		}
	}
	return mx
}
