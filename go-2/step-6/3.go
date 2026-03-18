package main

/*Дан сервер, доступный по адресу localhost:8082. По запросу localhost:8082/mark?name=<имя студента> сервер возвращает:

код 200 и значение оценки студента, если всё прошло успешно
код 404, если студент не найден
код 500, если у сервера проблема
Создайте функцию BestStudents(names []string) (string, error), которая выводит список студентов с оценками выше средней успеваемости из списка names в алфавитном порядке через запятую. Функция возвращает ошибку, если невозможно получить оценку хотя бы одного студента.*/
import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func Average(names []string) (int, error) {
	var f = false
	var summ = 0
	var count = 0

	for _, name := range names {

		tmpname := name
		resp, err := http.Get(fmt.Sprintf("http://localhost:8082/mark?name=%s", tmpname))
		if resp.StatusCode == 404 || resp.StatusCode == 500 || err != nil {
			f = true
			break
		}
		body, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		tmp, _ := strconv.Atoi(string(body))
		summ += tmp
		count++

	}

	if f {
		return 0, fmt.Errorf("There is no one rating")
	}
	return summ / count, nil
}
func GetMark(name string) int {
	resp, _ := http.Get(fmt.Sprintf("http://localhost:8082/mark?name=%s", name))
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	tmp, _ := strconv.Atoi(string(body))
	return tmp
}
func BestStudents(names []string) (string, error) {
	avg, err := Average(names)
	best := []string{}
	if err != nil {
		return "", err
	}
	for _, name := range names {
		if GetMark(name) > avg {
			best = append(best, name)
		}
	}
	return strings.Join(best, ","), nil

}
