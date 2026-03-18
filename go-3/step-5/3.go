package main

/*Реализуйте HTTP-сервер на Go, который:

При запросе по адресу / возвращает следующее число Фибоначчи (0, 1, 1, 2, 3, ...).
При запросе по адресу /metrics возвращает строку вида rpc_duration_milliseconds_count N, где N — число запросов к хендлеру /.
Требования:

Подсчёт числа запросов реализовать через middleware Metrics(next http.HandlerFunc).
Middleware Metrics должен увеличивать метрику только для обработчика /, но не для /metrics.
Используйте мьютекс для синхронизации доступа к переменным состояния.
Число запросов не должно увеличиваться при обращении к /metrics.*/
import (
	"fmt"
	"net/http"
	"sync"
)

var (
	count = 0
	mx    = &sync.Mutex{}
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

func Metrics(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mx.Lock()
		count++
		mx.Unlock()

		next.ServeHTTP(w, r)

	}
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	mx.Lock()
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", count)
	mx.Unlock()
}

func FibHandler(w http.ResponseWriter, r *http.Request) {
	mx.Lock()
	fmt.Fprintln(w, Fib(count))
	mx.Unlock()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/metrics", MetricsHandler)
	mux.Handle("/", Metrics(FibHandler))

	http.ListenAndServe(":8080", mux)
}
