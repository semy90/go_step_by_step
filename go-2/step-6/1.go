package main

/*Дан сервер, доступный по адресу localhost:8082. По запросу localhost:8082/mark?name=<имя студента> сервер возвращает:

код 200 и значение оценки студента, если всё прошло успешно
код 404, если студент не найден
код 500, если у сервера проблема
Напишите функцию Compare(name1, name2 string) (string, error), которая сравнивает оценки двух студентов с именами name1 и name2 и выводит > (оценка студента 1 больше оценки студента 2), < (оценка студента 1 меньше оценки студента 2) или = (оценка студента 1 равна оценке студента 2). Функция возвращает ошибку, если невозможно получить оценку хотя бы одного студента.*/
import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Response struct {
	Name   string
	Grade  int
	Status int
	Err    error
}

func Compare(name1, name2 string) (string, error) {
	var f = false
	std := []string{name1, name2}
	students := []Response{}
	for i := 0; i < 2; i++ {
		resp, err := http.Get(fmt.Sprintf("http://localhost:8082/mark?name=%s", std[i]))
		if resp.StatusCode == 500 || resp.StatusCode == 404 {
			students = append(students, Response{Name: std[i], Status: resp.StatusCode, Err: err})
			f = true
			continue
		}
		defer resp.Body.Close()
		bodyBytes, _ := io.ReadAll(resp.Body)
		n, _ := strconv.Atoi(string(bodyBytes))

		students = append(students, Response{Name: std[i], Status: resp.StatusCode, Err: nil, Grade: n})

	}
	if f {
		return "", fmt.Errorf("There is no one rating")
	}
	if students[0].Grade == students[1].Grade {
		return "=", nil
	}
	if students[0].Grade > students[1].Grade {
		return ">", nil
	}
	return "<", nil
}
