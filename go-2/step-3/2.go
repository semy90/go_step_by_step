package main

/*
Напишите функцию

func Receive(ch chan int) int
которая возвращает число, полученное из канала ch.
*/
func Receive(ch chan int) int {
	return <-ch
}
