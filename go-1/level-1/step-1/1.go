package main

/*
Напишите программу, которая принимает на вход строку и, если это строка Go, выводит Go!, а если нет, — Я знаю только Go!.
*/
import (
	"fmt"
)

func main() {
	var lang string
	fmt.Scanln(&lang)
	if lang == "Go" {
		fmt.Println("Go!")
	} else {
		fmt.Println("Я знаю только Go!")
	}
}
