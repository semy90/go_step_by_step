package main

/*
Напишите универсальную функцию Filter

func Filter[T any](arr []T, predicate func(T) bool) []T
которая фильтрует слайс данных по условию predicate
*/
func Filter[T any](arr []T, predicate func(T) bool) []T {
	answ := []T{}
	for _, el := range arr {
		if predicate(el) {
			answ = append(answ, el)
		}
	}
	return answ
}
