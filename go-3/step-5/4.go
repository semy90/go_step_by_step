package main

/*Реализуйте веб-сервер на Go, который защищает доступ к ресурсу с помощью базовой HTTP-авторизации.

Хендлер должен обрабатывать запросы по пути /answer/.
Для доступа к ресурсу клиент должен передавать заголовок Authorization с логином и паролем в формате Basic Auth.
В рамках задачи имя пользователя и пароль могут быть любыми, но оба параметра должны быть НЕ пустыми.
Если авторизация прошла успешно, сервер возвращает ответ:
Welcome, <username>!
где <username> — имя пользователя, указанное в запросе.
Если заголовок Authorization отсутствует, или логин или пароль пустые, или формат неверный, сервер возвращает статус 401 Unauthorized и устанавливает заголовок:
WWW-Authenticate: Basic realm="Restricted"
Валидацию логина и пароля вынесите в middleware-функцию Authorization(http.HandlerFunc).
Пример успешного запроса:

curl -u john:secret http://localhost:8080/answer/
# Welcome, john!
Пример неуспешного запроса:

curl http://localhost:8080/answer/
# HTTP/1.1 401 Unauthorized
# WWW-Authenticate: Basic realm="Restricted"*/
import (
	"fmt"
	"net/http"
)

func answerHandler(w http.ResponseWriter, r *http.Request) {
	name, _, _ := r.BasicAuth()
	fmt.Fprintf(w, "Welcome, %s!", name)
}
func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name, pass, ok := r.BasicAuth()
		if name == "" || pass == "" || !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/answer/", Authorization(answerHandler))

	http.ListenAndServe(":8080", mux)
}
