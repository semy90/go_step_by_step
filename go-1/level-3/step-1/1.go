package main

/*
рограммист Арсений хочет оптимизировать свою программу. Поскольку он много работает со строками, ему нужно узнать, сколько байтов они занимают и сколько символов содержат.
Напишите функцию CountLengthAndBytes(first, second string) string, которая считает общее количество символов (рун) в двух строках и их суммарное количество байт.
На выход программа должна вернуть строку: Объединённая строка: *объединённая строка из двух входных строк*. Количество байт: *число байт*. Количество символов: *число символов*.
*/

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func CountLengthAndBytes(first, second string) string {
	mergedString := strings.Join([]string{first, second}, "")
	bits1 := len(mergedString)
	lenght := utf8.RuneCountInString(mergedString)
	ans := fmt.Sprintf("Объединённая строка: %s. Количество байт: %d. Количество символов: %d.", mergedString, bits1, lenght)
	return ans
}
