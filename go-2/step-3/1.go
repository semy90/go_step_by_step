package main

/*
Напишите программу с функцией

func Send(ch chan int, num int)
которая отправляет число num в канал ch.
*/
func Send(ch chan int, num int) {
	go func() {
		ch <- num
	}()
}
