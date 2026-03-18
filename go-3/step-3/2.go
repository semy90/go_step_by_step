package main

/*Реализуйте функцию BuildHTTPResponse(statusLine, headers, body string) string,
которая формирует строку простого HTTP-ответа из следующих параметров:

statusLine — строка статуса (например, "HTTP/1.1 200 OK")
headers — дополнительные заголовки (каждый заголовок заканчивается символом новой строки, например, "Content-Type: application/json\r\nDate: Mon, 12 Apr 2021 10:00:00 GMT\r\n")
body — тело ответа (может быть пустым)
Функция должна вернуть корректную строку HTTP-ответа, в которой:

первая строка — это статус ответа (например, "HTTP/1.1 200 OK")
далее идут заголовки (если есть)
затем пустая строка (разделяет заголовки и тело)
и далее тело ответа (если оно есть)
*/
import (
	"strings"
)

func BuildHTTPResponse(statusLine, headers, body string) string {
	var response strings.Builder

	response.WriteString(statusLine)
	response.WriteString("\r\n")

	if headers != "" {
		cleanHeaders := strings.TrimRight(headers, "\r\n")
		response.WriteString(cleanHeaders)
		response.WriteString("\r\n")
	}

	response.WriteString("\r\n")

	if body != "" {
		response.WriteString(body)
	}

	return response.String()
}
