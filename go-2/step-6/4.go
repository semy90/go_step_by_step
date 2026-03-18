package main

/*Дан сервер доступный по адресу localhost:8082. По запросу localhost:8082/mark?name=<имя студента> сервер возвращает:

код 200 и значение оценки студента, если всё прошло успешно
код 404, если студент не найден
код 500, если у сервера проблема
Напишите функцию CompareList(names []string) (map[string]string, error), которая выводит карту. Её ключ — имя студента из списка name, а значение — > (оценка студента больше средней оценки студентов), < (оценка студента меньше средней оценки студентов) или = (оценка студента равна средней оценки студентов). Функция возвращает ошибку, если невозможно получить оценку хотя бы одного студента.
*/
import (
	"fmt"
	"io"
	"net/http"
	"strconv"
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
func CompareList(names []string) (map[string]string, error) {
	avg, err := Average(names)
	mp := make(map[string]string)
	if err != nil {
		return map[string]string{}, err
	}
	for _, name := range names {
		mark := GetMark(name)
		if mark > avg {
			mp[name] = ">"
		} else if mark == avg {
			mp[name] = "="
		} else {
			mp[name] = "<"
		}
	}
	return mp, err
}
