package main

/*Напишите веб-сервер, который будет считать метрики времени ответа сервиса. Возьмите в качестве основы предыдущий веб-сервер из предыдущего задания и добавьте к нему хендлер /metrics — он будет возвращать число запросов, которые сделаны за числом Фибоначчи:*/
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
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%d", k)
	})
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
