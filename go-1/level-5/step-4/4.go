package main

/*Иногда так хочется изучить новую сортировку... Вот и программист Арсений на выходных решил потратить на это время.

Даны два слайса. Напишите программу с функцией SortAndMerge(left, right []int) []int, которая объединит слайсы в один отсортированный в два этапа:

отсортировать каждый слайс
объединить полученные слайсы в один
Кстати, именно так работает алгоритм сортировки слиянием (merge sort).*/
import (
	"sort"
)

func SortAndMerge(left, right []int) []int {
	ans := []int{}
	for _, el := range left {
		ans = append(ans, el)
	}
	for _, el := range right {
		ans = append(ans, el)
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	return ans
}
