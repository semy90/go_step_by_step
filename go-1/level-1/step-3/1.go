package main

/*
Доработайте задачу «Go или не Go, вот в чём вопрос»: напишите программу, которая принимает на вход 5 строк и, если это строка Go, выводит: Go!, а если нет, то: Я знаю только Go!.
*/
import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		var s string
		fmt.Scanln(&s)
		if s == "Go" {
			fmt.Println("Go!")
		} else {
			fmt.Println("Я знаю только Go!")
		}
	}
}
