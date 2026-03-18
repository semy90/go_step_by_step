package main

/*
Напишите веб-сервер, который будет возвращать приветствие с именем пользователя, полученным из параметра запроса. Если параметр пустой или отсутствует, сервер должен вернуть приветствие "hello stranger". Если в ответе не только английские буквы, приветствие должно гласить: "hello dirty hacker" (можете применить функцию из пакета strings или regexp.Match). Веб-сервер должен отвечать на порт :8080.
*/
import (
	"fmt"
	"net/http"
)

func isURLValid(rawURL string) bool {
	for _, el := range rawURL {
		if !(('a' <= el && el <= 'z') || ('A' <= el && el <= 'Z') || el == '/' || el == ':' || el == '.' || el == '&' || el == '?') {
			return false
		}
	}
	return true
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if !isURLValid(r.URL.Path) {
			fmt.Fprintf(w, "hello dirty hacker")
			return
		}
		if name == "" {
			fmt.Fprintf(w, "hello stranger")
			return
		}
		fmt.Fprintf(w, "hello %s", name)
		return
	})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
