package main

/*
Да, если результаты двух задач идут друг за другом, соединять их легко. Но бывают ситуации и посложнее…
Дан слайс nums из 2n элементов в формате [x0,x1,...,xn,y0,y1,...,yn]. Создайте функцию Mix(nums []int) []int, которая вернёт слайс со значениями в таком порядке: [x0,y0,x1,y1,...,xn,yn].
*/
func Mix(nums []int) []int {
	ans := make([]int, len(nums))
	l1, l2 := 0, len(nums)/2

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			ans[i] = nums[l1]
			l1++
		} else {
			ans[i] = nums[l2]
			l2++
		}
	}
	return ans
}
