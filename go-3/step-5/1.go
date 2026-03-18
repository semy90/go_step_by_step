package main

/*Реализуйте HTTP-сервер на Go, который:

Обрабатывает запросы по пути /hello и возвращает строку Hello, middleware!.
Использует middleware-функцию Logger(next http.Handler) http.Handler, которая логирует сообщение incoming request, HTTP-метод и путь каждого запроса через пакет slog перед передачей управления следующему обработчику.
Пример работы:

curl -X POST http://localhost:8080/hello
В логах сервера при таком запросе должна появиться строка вида:

level=INFO msg="incoming request" method=POST path=/hello*/
import (
	"fmt"
	"log/slog"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("incoming request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, middleware!")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	handler := Logger(mux)

	http.ListenAndServe(":8080", handler)
}
