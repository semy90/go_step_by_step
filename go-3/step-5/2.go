package main

/*Реализуйте HTTP-сервер на Go, который:

Возвращает приветствие с именем пользователя, полученным из параметра name в строке запроса (/hello?name=John → hello John).
Если параметр name отсутствует или пустой, сервер должен вернуть приветствие "hello stranger".
Если параметр name содержит не только английские буквы (A-Z, a-z), сервер должен вернуть "hello dirty hacker".
Требования:

Валидация значения name на корректность (только буквы) должна происходить в middleware Sanitize(next http.HandlerFunc).
Подстановка значения "stranger" при отсутствии или пустом параметре должна происходить в middleware SetDefaultName(next http.HandlerFunc).*/
import (
	"fmt"
	"net/http"
)

func isNameValid(name string) bool {
	if name == "" {
		return true
	}
	for _, el := range name {
		if !(('a' <= el && el <= 'z') || ('A' <= el && el <= 'Z') || el == '/' || el == ':' || el == '.' || el == '&' || el == '?') {
			return false
		}
	}
	return true
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "hello %s", name)
	return
}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if isNameValid(name) {
			next.ServeHTTP(w, r)
			return
		}
		fmt.Fprintf(w, "hello dirty hacker")
		return
	}
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name != "" {
			next.ServeHTTP(w, r)

			return
		}
		fmt.Fprintf(w, "hello stranger")
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/hello", Sanitize(SetDefaultName(HelloHandler)))
	http.ListenAndServe(":8080", mux)
}
