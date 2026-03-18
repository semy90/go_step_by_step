package main

/*Напишите веб-сервер, который будет принимать запросы на получение следующего числа Фибоначчи и возвращать его значение.

Примеры запросов и ответов:

curl http://localhost:8080/

# 0

curl http://localhost:8080/

# 1

curl http://localhost:8080/

# 1

curl http://localhost:8080/

# 2
Сервер не сохраняет своё состояние между перезапусками: если закрыть программу и запустить её заново, подсчёт начнётся с 0.

Числа Фибоначчи — это последовательность чисел, где каждое следующее равно сумме двух предыдущих. Например, первые несколько чисел в ряду могут выглядеть так: 0, 1, 1, 2, 3, 5, 8, 13. Эта последовательность часто встречается в математике, науке и программировании. В вашей программе вы будете вычислять числа Фибоначчи и возвращать их пользователю по запросу через веб-сервер.*/
import (
	"fmt"
	"net/http"
)

func Fib(num int) int {
	n1, n2 := 0, 1
	if num == 1 {
		return n1
	}
	if num == 2 {
		return n2
	}
	k := 2
	for ; k <= num; k++ {
		tmp := n1
		n1 = n2
		n2 = tmp + n1
	}
	return n1
}
func main() {
	k := 1
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%d", Fib(k))
		fmt.Println(k, Fib(k))
		k += 1

	})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}
