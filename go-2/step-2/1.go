package main

/*Напишите программу с функцией

func ReadContent(filename string) string
которая принимает на вход путь к файлу, а возвращает его содержимое.

Если возникнет любая ошибка, возвращайте пустую строку.*/
import "os"

func ReadContent(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}
