package main

/*
Возьмём на доработку уже полюбившуюся вам задачу «Go или не Go, вот в чём вопрос». Напишите процедуру GoOrNot(), которая принимает на вход строку и, если это строка Go, выводит Go!, а если нет, то Я знаю только Go!
*/
import "fmt"

func GoOrNot(input string) {
	if input == "Go" {
		fmt.Println("Go!")
	} else {
		fmt.Println("Я знаю только Go!")
	}
}
