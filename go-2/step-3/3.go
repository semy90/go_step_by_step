package main

/*
Напишите функцию

func Send(ch1, ch2 chan int)
которая отправляет числа 0,1,2 в каналы ch1 и ch2 в отдельных горутинах.
*/
func Send(ch1, ch2 chan int) {
	go func() {
		ch1 <- 0
		ch1 <- 1
		ch1 <- 2
	}()
	go func() {
		ch2 <- 0
		ch2 <- 1
		ch2 <- 2
	}()
}
