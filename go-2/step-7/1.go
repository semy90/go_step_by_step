package main

/*Сгенерируйте N-е число Фибоначчи. Установите ограничение на время вычисления. Если число будет вычислено в течение указанного времени, верните число Фибоначчи, в противном случае выдайте ошибку. Сигнатура функции:

func TimeoutFibonacci(n int, timeout time.Duration) (int, error)
n — номер числа timeout — время на операцию*/
import (
	"fmt"
	"time"
)

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	arr := []int{0, 1}
	for i := 2; i <= n; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
	}
	return arr[n]
}

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("n must be non-negative")
	}
	ch := make(chan int)
	go func() {
		answ := Fibonacci(n)
		ch <- answ
	}()
	select {
	case answ := <-ch:
		return answ, nil
	case <-time.After(timeout):
		return 0, fmt.Errorf("timeout")
	}
}
