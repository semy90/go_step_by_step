package main

/*Реализуйте функцию:

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int)
которая будет генерировать простые числа и писать их в канал prime_nums. C помощью функции time.AfterFunc остановите генерацию принудительно через 0.1 секунды

stop — канал для остановки генерации prime_nums — канал для вывода простых чисел N — число, до которого нужно генерировать числа*/
import (
	"time"
)

func IsPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	defer close(prime_nums)
	time.AfterFunc(time.Second/10, func() {
		close(stop)
	})
	for i := 2; i <= N; i++ {
		select {
		case <-stop:
			return
		default:
			if IsPrime(i) {
				prime_nums <- i
			}
		}
	}

}
