package main

/*Реализуйте функцию BuildHTTPRequest(method, url, host, headers, body string) string,
которая собирает строку простого HTTP-запроса из переданных параметров:

method — HTTP-метод (например, "GET", "POST")
url — путь к ресурсу (например, "/api/data")
host — заголовок Host (например, "example.com")
headers — дополнительные заголовки (каждый заголовок заканчивается символом новой строки, например, "Content-Type: application/json\r\nAuthorization: Bearer xyz\r\n")
body — тело запроса (может быть пустым).
Функция должна вернуть корректную строку HTTP-запроса, в которой:

первая строка — "<METHOD> <URL> HTTP/1.1"
далее строка "Host: <host>"
далее идут дополнительные заголовки (если есть)
затем пустая строка (разделяет заголовки и тело)
и далее тело запроса (если оно есть)*/
import (
	"fmt"
	"strings"
)

func BuildHTTPRequest(method, url, host, headers, body string) string {
	var request strings.Builder

	request.WriteString(fmt.Sprintf("%s %s HTTP/1.1\r\n", method, url))

	request.WriteString(fmt.Sprintf("Host: %s\r\n", host))

	if headers != "" {
		cleanHeaders := strings.TrimRight(headers, "\r\n")
		request.WriteString(cleanHeaders)
		request.WriteString("\r\n")
	}

	request.WriteString("\r\n")

	if body != "" {
		request.WriteString(body)
	}

	return request.String()
}
