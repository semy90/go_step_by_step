package main

/*
Программист Арсений любит ASCII и терпеть не может другие виды кодировок.

Напишите функцию CheckOnlyASCII(s string) bool, которая проверяет, символы какой кодировки содержатся в строке, и возвращает true, если все они в ASCII, а false, если нет.
*/
import (
	"unicode"
	"unicode/utf8"
)

func CheckOnlyASCII(s string) bool {
	for i := 0; i < utf8.RuneCountInString(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
