package main

/*
Напишите функцию

func Process(nums []int) chan int
которая создаёт буферизованный канал (capacity = 10) и записывает в него числа из nums без создания горутин. Функция должна вернуть созданный канал.
*/
func Process(nums []int) chan int {
	ch := make(chan int, 10)
	for _, el := range nums {
		ch <- el
	}
	return ch
}
