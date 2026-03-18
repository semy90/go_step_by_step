package main

/*Реализуйте функцию ParseHTTPStatus(statusLine string) (code int, reason string),
которая принимает строку статуса HTTP-ответа (например, "HTTP/1.1 404 Not Found")
и возвращает числовой код статуса и текстовое пояснение (reason phrase).

Гарантируется, что на вход всегда подается строка вида:
"HTTP/<версия> <код> <reason phrase>"
Код всегда трехзначный, reason phrase может содержать пробелы.
Примечания
Пример:

Вызов:

ParseHTTPStatus("HTTP/1.1 418 I'm a teapot")
Должен вернуть:

418, "I'm a teapot"*/
import (
	"net/http"
	"strconv"
	"strings"
)

func ParseHTTPStatus(statusLine string) (code int, reason string) {
	ans := strings.Split(statusLine, " ")
	cd, _ := strconv.Atoi(ans[1])
	return cd, http.StatusText(cd)
}
