package main

/*Реализуйте функцию MakeCurlCommand(method, url, headers, body string) string,
которая формирует строку команды для утилиты curl на основе переданных параметров:

method — HTTP-метод (например, "GET", "POST"). Для "GET" параметр -X можно не указывать
url — полный адрес (например, "https://example.com/api/data")
headers — дополнительные заголовки (каждый заголовок заканчивается символом новой строки, например, "Content-Type: application/json\nAuthorization: Bearer xyz\n")
body — тело запроса (использовать только если оно не пустое; передавать через --data '...')
Функция должна вернуть строку команды, которую можно выполнить в терминале.
Каждый заголовок должен указываться через -H '...', тело через --data '...',
URL — последний параметр. Заголовки и тело должны корректно экранироваться одинарными кавычками.*/
import (
	"fmt"
	"strings"
)

func MakeCurlCommand(method, url, headers, body string) string {
	var curlComand strings.Builder
	curlComand.WriteString("curl ")
	if method != "GET" {
		curlComand.WriteString(fmt.Sprintf("-X %s ", method))
	}

	if headers != "" {
		s := strings.Split(headers, "\n")
		for i, header := range s {
			if i != len(s)-1 {
				curlComand.WriteString(fmt.Sprintf("-H '%s' ", header))
			}
		}
	}
	if body != "" {
		curlComand.WriteString(fmt.Sprintf("--data '%s' ", body))
	}
	curlComand.WriteString(url)
	return curlComand.String()
}
