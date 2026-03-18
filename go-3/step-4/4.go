package main

/*еализуйте HTTP-сервер с помощью http.HandleFunc, который по запросу на адрес /echo возвращает значение параметра msg из строки запроса.
Если параметр msg отсутствует или пустой — вернуть текст "empty".

Примеры:

curl "http://localhost:8080/echo?msg=hello"
# hello

curl "http://localhost:8080/echo?msg="
# empty

curl "http://localhost:8080/echo"
# empty*/
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		msg := r.URL.Query().Get("msg")
		if msg == "" {
			fmt.Fprintf(w, "empty")
		}
		fmt.Fprintf(w, msg)
	})
	http.ListenAndServe(":8080", nil)
}
